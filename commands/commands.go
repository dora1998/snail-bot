package commands

import (
	"fmt"
	"github.com/dora1998/snail-bot/repository"
	"regexp"
	"time"
)

type Command struct {
	Name       string
	HandleFunc func(body string, username string, statusId int64, repo Repository)
}

type CommandHandler struct {
	commands   []*Command
	repository Repository
}

type Repository interface {
	Add(body string, deadline time.Time, createdBy string) *repository.Task
	Remove(id string) error
	GetAllTasks() []repository.Task
	GetTaskById(id string) *repository.Task
	GetTaskByBody(body string) *repository.Task
}

var CmdHandler = newCommandHandler()

func newCommandHandler() *CommandHandler {
	return &CommandHandler{commands: []*Command{}}
}

func SetRepository(repo Repository) {
	CmdHandler.repository = repo
}

func (h *CommandHandler) Resolve(text string, username string, statusId int64) error {
	if h.repository == nil {
		return fmt.Errorf("repo is not set")
	}

	regexpObj := regexp.MustCompile("^(\\S+)(\\s(.+))*$")
	res := regexpObj.FindStringSubmatch(text)
	if res == nil {
		return fmt.Errorf("failed resolve (incorrect pattern)")
	}

	commandName, commandBody := res[1], res[3]
	fmt.Printf("%s: %s\n", commandName, commandBody)
	for _, c := range h.commands {
		if commandName == c.Name {
			c.HandleFunc(commandBody, username, statusId, h.repository)
			return nil
		}
	}

	return fmt.Errorf("failed resolve (no match)")
}

func (h *CommandHandler) AddCommand(c *Command) {
	h.commands = append(h.commands, c)
}
