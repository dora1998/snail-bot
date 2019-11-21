package commands

import (
	"fmt"
)

func (h *CommandHandler) list(username string, statusId int64) {
	fmt.Printf("list (%v)\n", statusId)

	output := ""
	for _, t := range h.repository.GetAllTasks() {
		output += fmt.Sprintf("[%s〆]%s\n", t.Deadline.Format("01/02"), t.Body)
	}

	if output == "" {
		output = "現在出ている課題はありません🎉"
	}

	_, err := h.twitterClient.Reply(output, statusId)
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
}
