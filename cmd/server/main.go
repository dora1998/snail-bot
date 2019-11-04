package main

import (
	"encoding/json"
	"fmt"
	"github.com/dora1998/snail-bot/db"
	"github.com/dora1998/snail-bot/repository"
	"github.com/dora1998/snail-bot/utils"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type CallbackBody struct {
	Text        string `json:"text"`
	UserName    string `json:"user_name"`
	LinkToTweet string `json:"link_to_tweet"`
	CreatedAt   string `json:"created_at"`
}

func main() {
	dbConfig, err := utils.ReadDBConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbInstance, err := db.NewDBInstance(dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.RunMigration(dbInstance)
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = dbInstance.Close()

	handler := utils.NewCommandHandler()
	handler.AddCommand(&utils.Command{
		Name: "追加",
		HandleFunc: func(body string, username string, statusId int64) {
			defer dbInstance.Close()

			fmt.Printf("add: %s (%v)\n", body, statusId)

			regexpObj := regexp.MustCompile("^(.+)\\s([0-9]+/[0-9]+)$")
			parsedBody := regexpObj.FindStringSubmatch(body)
			if parsedBody == nil {
				fmt.Printf("ParseError: %#v\n", body)
				return
			}

			parsedDate, err := utils.ParseDateStr(parsedBody[2])
			if err != nil {
				fmt.Printf("ParseDateError: %#v\n", parsedBody[2])
				return
			}

			dbInstance, err := db.NewDBInstance(dbConfig)
			if err != nil {
				log.Fatal(err.Error())
			}
			repo := repository.NewDBRepository(dbInstance)

			task := repo.Add(parsedBody[1], parsedDate, username)
			fmt.Printf("added: %#v\n", task)
			client := utils.NewTwitterClient()
			client.Reply(fmt.Sprintf("%v\n%#v", parsedBody[1], parsedBody[2]), statusId)
		},
	})
	handler.AddCommand(&utils.Command{
		Name: "一覧",
		HandleFunc: func(_ string, username string, statusId int64) {
			defer dbInstance.Close()

			fmt.Printf("list (%v)\n", statusId)

			dbInstance, err := db.NewDBInstance(dbConfig)
			if err != nil {
				log.Fatal(err.Error())
			}
			repo := repository.NewDBRepository(dbInstance)

			output := ""
			for _, t := range repo.GetAllTasks() {
				output += fmt.Sprintf("%s(%s)\n", t.Body, t.Deadline.Format("1/2"))
			}

			client := utils.NewTwitterClient()
			client.Reply(output, statusId)
		},
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		callbackBody := &CallbackBody{}
		err := json.NewDecoder(r.Body).Decode(callbackBody)
		if err != nil {
			fmt.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Printf("%#v\n", callbackBody)
		statusId, err := utils.ExtractStatusIdFromUrl(callbackBody.LinkToTweet)
		if err != nil {
			fmt.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		text, err := utils.ExtractBody(callbackBody.Text)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		err = handler.Resolve(text, callbackBody.UserName, statusId)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	})

	http.HandleFunc("/exec", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		err = handler.Resolve(string(body), "exec", 1185595390327787520) // statusId is dummy
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}
