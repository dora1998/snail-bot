package commands

import (
	"fmt"
	"github.com/dora1998/snail-bot/utils"
	"regexp"
)

func (h *CommandHandler) add(body string, username string, statusId int64) {
	fmt.Printf("add: %s (%v)\n", body, statusId)

	if !h.twitterClient.IsFollwing(username) {
		fmt.Printf("PermissionError: not following @%v\n", username)
		h.twitterClient.Reply("この操作はフォローされている人しかできません🙇‍♂️", statusId)
		return
	}

	regexpObj := regexp.MustCompile("^(.+)\\s([0-9]+/[0-9]+)$")
	parsedBody := regexpObj.FindStringSubmatch(body)
	if parsedBody == nil {
		fmt.Printf("ParseError: %#v\n", body)
		h.twitterClient.Reply("タスクの追加に失敗しました…", statusId)
		return
	}

	parsedDate, err := utils.ParseDateStr(parsedBody[2])
	if err != nil {
		fmt.Printf("ParseDateError: %#v\n", parsedBody[2])
		h.twitterClient.Reply("タスクの追加に失敗しました…", statusId)
		return
	}

	task := h.repository.Add(parsedBody[1], parsedDate, username)
	if task == nil {
		fmt.Printf("DatabaseError: %v\n", body)
		h.twitterClient.Reply("タスクの追加に失敗しました…", statusId)
		return
	}

	fmt.Printf("added: %#v\n", task)
	h.twitterClient.Reply(fmt.Sprintf("%v (%v)\n", parsedBody[1], parsedBody[2]), statusId)
}
