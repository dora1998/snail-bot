package commands

import (
	"fmt"
)

func (h *CommandHandler) remove(body string, username string, statusId int64) {
	fmt.Printf("remove: %s (%v)\n", body, statusId)

	if !h.twitterClient.IsFollowing(username) {
		fmt.Printf("PermissionError: not following @%v\n", username)
		_, err := h.twitterClient.Reply("この操作はフォローされている人しかできません🙇‍♂️", statusId)
		if err != nil {
			_ = fmt.Errorf(err.Error())
		}
		return
	}

	task := h.repository.GetTaskByBody(body)
	if task == nil {
		fmt.Printf("TaskNotFound: %v\n", body)
		_, err := h.twitterClient.Reply("該当するタスクが見つかりません", statusId)
		if err != nil {
			_ = fmt.Errorf(err.Error())
		}
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
