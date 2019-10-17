package main

import (
	"fmt"
	"regexp"
)

type Command struct {
	name       string
	handleFunc func(body string, username string, statusId int64)
}

type CommandHandler struct {
	commands []*Command
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{commands: []*Command{}}
}

func (h *CommandHandler) resolve(text string, username string, statusId int64) error {
	regexpObj := regexp.MustCompile("^(\\S+)(\\s(.+))*$")
	res := regexpObj.FindStringSubmatch(text)
	if res == nil {
		return fmt.Errorf("failed resolve (incorrect pattern)")
	}

	commandName, commandBody := res[1], res[3]
	fmt.Printf("%s: %s\n", commandName, commandBody)
	for _, c := range h.commands {
		if commandName == c.name {
			c.handleFunc(commandBody, username, statusId)
			return nil
		}
	}

	return fmt.Errorf("failed resolve (no match)")
}

func (h *CommandHandler) addCommand(c *Command) {
	h.commands = append(h.commands, c)
}
