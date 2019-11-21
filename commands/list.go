package commands

import (
	"fmt"
)

func (h *CommandHandler) list(username string, statusId int64) {
	fmt.Printf("list (%v)\n", statusId)

	output := ""
	for _, t := range h.repository.GetAllTasks() {
		output += fmt.Sprintf("・%s【%s〆】\n", t.Body, t.Deadline.Format("1/2"))
	}

	if output == "" {
		output = "現在出ている課題はありません🎉"
	}

	_, err := h.twitterClient.Reply(output, statusId)
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
}
