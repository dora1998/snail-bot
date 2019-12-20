package commands

import (
	"fmt"
	"log"
)

func (h *CommandHandler) list(statusId int64) {
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
		log.Fatal(err.Error())
	}
}
