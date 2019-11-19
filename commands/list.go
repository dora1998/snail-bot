package commands

import (
	"fmt"
	"github.com/dora1998/snail-bot/utils"
)

func (h *CommandHandler) list(username string, statusId int64) {
	fmt.Printf("list (%v)\n", statusId)

	output := ""
	for _, t := range h.repository.GetAllTasks() {
		output += fmt.Sprintf("%s(%s)\n", t.Body, t.Deadline.Format("1/2"))
	}

	if output == "" {
		output = "現在出ている課題はありません🎉"
	}

	client := utils.NewTwitterClient()
	client.Reply(output, statusId)
}
