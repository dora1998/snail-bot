package commands

import (
	"fmt"
	"github.com/dora1998/snail-bot/utils"
)

var cmdList = &Command{
	Name: "一覧",
	HandleFunc: func(_ string, username string, statusId int64, repo Repository) {
		fmt.Printf("list (%v)\n", statusId)

		output := ""
		for _, t := range repo.GetAllTasks() {
			output += fmt.Sprintf("%s(%s)\n", t.Body, t.Deadline.Format("1/2"))
		}

		if output == "" {
			output = "現在出ている課題はありません🎉"
		}

		client := utils.NewTwitterClient()
		client.Reply(output, statusId)
	},
}

func init() {
	CmdHandler.AddCommand(cmdList)
}
