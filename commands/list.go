package commands

import (
	"fmt"
	"github.com/dora1998/snail-bot/repository"
	"github.com/dora1998/snail-bot/utils"
	"github.com/jmoiron/sqlx"
)

var cmdList = &Command{
	Name: "一覧",
	HandleFunc: func(_ string, username string, statusId int64, db *sqlx.DB) {
		fmt.Printf("list (%v)\n", statusId)

		repo := repository.NewDBRepository(db)

		output := ""
		for _, t := range repo.GetAllTasks() {
			output += fmt.Sprintf("%s(%s)\n", t.Body, t.Deadline.Format("1/2"))
		}

		client := utils.NewTwitterClient()
		client.Reply(output, statusId)
	},
}

func init() {
	CmdHandler.AddCommand(cmdList)
}
