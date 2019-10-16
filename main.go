package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CallbackBody struct {
	Text        string `json:"text"`
	UserName    string `json:"user_name"`
	LinkToTweet string `json:"link_to_tweet"`
	CreatedAt   string `json:"created_at"`
}

func main() {
	repo, _ := NewTaskRepository()

	handler := NewCommandHandler()
	handler.addCommand(&Command{
		name: "追加",
		handleFunc: func(body string, statusId int64) {
			fmt.Printf("add: %s (%v)\n", body, statusId)
			// TODO: 日付処理
			repo.Add(body, "10/16")
			client := NewTwitterClient()
			client.reply(body, statusId)
		},
	})
	handler.addCommand(&Command{
		name: "一覧",
		handleFunc: func(_ string, statusId int64) {
			fmt.Printf("list (%v)\n", statusId)

			output := ""
			for _, t := range repo.GetAllTasks() {
				output += t.Body + "\n"
			}

			client := NewTwitterClient()
			client.reply(output, statusId)
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
		statusId, err := extractStatusIdFromUrl(callbackBody.LinkToTweet)
		if err != nil {
			fmt.Printf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		text, err := extractBody(callbackBody.Text)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		err = handler.resolve(text, statusId)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}
