package commands

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"regexp"
)

type Command struct {
	Name       string
	HandleFunc func(body string, username string, statusId int64, db *sqlx.DB)
}

type CommandHandler struct {
	commands []*Command
	db       *sqlx.DB
}

var CmdHandler = NewCommandHandler()

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{commands: []*Command{}}
}

func (h *CommandHandler) SetDBInstance(db *sqlx.DB) {
	h.db = db
}

func (h *CommandHandler) Resolve(text string, username string, statusId int64) error {
	if h.db == nil {
		return fmt.Errorf("db instance is not set")
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
			c.HandleFunc(commandBody, username, statusId, h.db)
			return nil
		}
	}

	return fmt.Errorf("failed resolve (no match)")
}

func (h *CommandHandler) AddCommand(c *Command) {
	h.commands = append(h.commands, c)
}
