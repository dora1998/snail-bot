package commands

import (
	"fmt"
	"github.com/dora1998/snail-bot/utils"
)

var cmdRemove = &Command{
	Name: "削除",
	HandleFunc: func(body string, username string, statusId int64, repo Repository) {
		fmt.Printf("remove: %s (%v)\n", body, statusId)
		client := utils.NewTwitterClient()

		if !client.IsFollwing(username) {
			fmt.Printf("PermissionError: not following @%v\n", username)
			client.Reply("この操作はフォローされている人しかできません🙇‍♂️", statusId)
			return
		}

		task := repo.GetTaskByBody(body)
		if task == nil {
			fmt.Printf("TaskNotFound: %v\n", body)
			client.Reply("該当するタスクが見つかりません", statusId)
			return
		}

		err := repo.Remove(task.Id)
		if err != nil {
			fmt.Printf("DatabaseError: %#v\n", err)
			return
		}
		fmt.Printf("removed: %#v\n", task)

		err = client.CreateFavorite(statusId)
		if err != nil {
			fmt.Printf("FavoriteError: %#v\n", err)
		}
	},
}

func init() {
	CmdHandler.AddCommand(cmdRemove)
}
