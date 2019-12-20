package cmd

import (
	"fmt"
	"github.com/dora1998/snail-bot/http"
	"github.com/spf13/cobra"
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
		server := http.NewServer()
		err := server.Start()
		if err != nil {
			fmt.Printf(err.Error())
		}
	},
}
