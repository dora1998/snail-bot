package commands

import (
	"fmt"
	"github.com/dora1998/snail-bot/utils"
	"log"
	"regexp"
)

func (h *CommandHandler) add(body string, username string, statusId int64) {
	fmt.Printf("add: %s (%v)\n", body, statusId)

	if !h.twitterClient.IsFollowing(username) {
		fmt.Printf("PermissionError: not following @%v\n", username)
		_, err := h.twitterClient.Reply("この操作はフォローされている人しかできません🙇‍♂️", statusId)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	regexpObj := regexp.MustCompile("^(.+)\\s([0-9]+/[0-9]+)$")
	parsedBody := regexpObj.FindStringSubmatch(body)
	if parsedBody == nil {
		fmt.Printf("ParseError: %#v\n", body)
		_, err := h.twitterClient.Reply("タスクの追加に失敗しました…", statusId)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	parsedDate, err := utils.ParseDateStr(parsedBody[2])
	if err != nil {
		fmt.Printf("ParseDateError: %#v\n", parsedBody[2])
		_, err := h.twitterClient.Reply("タスクの追加に失敗しました…", statusId)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	task := h.repository.Add(parsedBody[1], parsedDate, username)
	if task == nil {
		fmt.Printf("DatabaseError: %v\n", body)
		_, err := h.twitterClient.Reply("タスクの追加に失敗しました…", statusId)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	fmt.Printf("added: %#v\n", task)
	_, err = h.twitterClient.Reply(fmt.Sprintf("タスクを追加しました！\n%v (%v)", parsedBody[1], parsedDate.Format("2006/1/2")), statusId)
	if err != nil {
		log.Fatal(err.Error())
	}
}
