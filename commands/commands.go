package commands

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dora1998/snail-bot/repository"
	"regexp"
	"time"
)

type Command struct {
	Name       string
	HandleFunc func(body string, username string, statusId int64, repo Repository)
}

type CommandHandler struct {
	repository    Repository
	twitterClient TwitterClient
}

type Repository interface {
	Add(body string, deadline time.Time, createdBy string) *repository.Task
	Remove(id string) error
	GetAllTasks() []repository.Task
	GetTaskById(id string) *repository.Task
	GetTaskByBody(body string) *repository.Task
}

type TwitterClient interface {
	Tweet(msg string) (*twitter.Tweet, error)
	Reply(msg string, tweetId int64) (*twitter.Tweet, error)
	CreateFavorite(tweetId int64) error
	IsFollowing(screenName string) bool
	TweetLongText(text string, headText string) ([]*twitter.Tweet, error)
}

func NewCommandHandler(repo Repository, twitterClient TwitterClient) *CommandHandler {
	return &CommandHandler{repository: repo, twitterClient: twitterClient}
}

func (h *CommandHandler) Resolve(text string, username string, statusId int64) error {
	regexpObj := regexp.MustCompile("^(\\S+)(\\s(.+))*$")
	res := regexpObj.FindStringSubmatch(text)
	if res == nil {
		return fmt.Errorf("failed resolve (incorrect pattern)")
	}

	commandName, commandBody := res[1], res[3]
	fmt.Printf("%s: %s\n", commandName, commandBody)

	switch commandName {
	case "追加":
		h.add(commandBody, username, statusId)
	case "削除":
		h.remove(commandBody, username, statusId)
	case "一覧":
		h.list(statusId)
	default:
		return fmt.Errorf("failed resolve (no match)")
	}
	return nil
}
