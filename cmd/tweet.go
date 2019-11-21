package cmd

import (
	"fmt"
	"github.com/dora1998/snail-bot/db"
	"github.com/dora1998/snail-bot/repository"
	"github.com/dora1998/snail-bot/twitter"
	"github.com/dora1998/snail-bot/utils"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func init() {
	rootCmd.AddCommand(tweetCmd)
}

var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Tweet current tasks",
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
		defer dbInstance.Close()

		repo := repository.NewDBRepository(dbInstance)

		output := ""
		for _, t := range repo.GetAllTasks() {
			output += fmt.Sprintf("[%s〆]%s\n", t.Deadline.Format("01/02"), t.Body)
		}

		client := twitter.NewTwitterClient()
		_, err = client.TweetLongText(output, fmt.Sprintf("🐌 出ている課題(%s) [{paged}/{pages}]", time.Now().Format("1/2")))
		if err != nil {
			_ = fmt.Errorf(err.Error())
		}
	},
}
