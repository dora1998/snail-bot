package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/dora1998/snail-bot/commands"
	"github.com/dora1998/snail-bot/db"
	"github.com/dora1998/snail-bot/repository"
	"github.com/dora1998/snail-bot/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

type CallbackBody struct {
	Text        string `json:"text"`
	UserName    string `json:"user_name"`
	LinkToTweet string `json:"link_to_tweet"`
	CreatedAt   string `json:"created_at"`
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run bot server",
	Run: func(cmd *cobra.Command, args []string) {
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

		repo := repository.NewDBRepository(dbInstance)
		twitterClient := utils.NewTwitterClient()
		handler := commands.NewCommandHandler(repo, twitterClient)
		defer dbInstance.Close()

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

		_ = http.ListenAndServe(":8080", nil)
	},
}