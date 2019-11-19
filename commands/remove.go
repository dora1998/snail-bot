package commands

import (
	"fmt"
)

func (h *CommandHandler) remove(body string, username string, statusId int64) {
	fmt.Printf("remove: %s (%v)\n", body, statusId)

	if !h.twitterClient.IsFollwing(username) {
		fmt.Printf("PermissionError: not following @%v\n", username)
		h.twitterClient.Reply("この操作はフォローされている人しかできません🙇‍♂️", statusId)
		return
	}

	task := h.repository.GetTaskByBody(body)
	if task == nil {
		fmt.Printf("TaskNotFound: %v\n", body)
		h.twitterClient.Reply("該当するタスクが見つかりません", statusId)
		return
	}

	err := h.repository.Remove(task.Id)
	if err != nil {
		fmt.Printf("DatabaseError: %#v\n", err)
		return
	}
	fmt.Printf("removed: %#v\n", task)

	err = h.twitterClient.CreateFavorite(statusId)
	if err != nil {
		fmt.Printf("FavoriteError: %#v\n", err)
	}
}
