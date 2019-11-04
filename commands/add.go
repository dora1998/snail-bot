package commands

import (
	"fmt"
	"github.com/dora1998/snail-bot/repository"
	"github.com/dora1998/snail-bot/utils"
	"github.com/jmoiron/sqlx"
	"regexp"
)

var cmdAdd = &Command{
	Name: "追加",
	HandleFunc: func(body string, username string, statusId int64, db *sqlx.DB) {
		fmt.Printf("add: %s (%v)\n", body, statusId)

		regexpObj := regexp.MustCompile("^(.+)\\s([0-9]+/[0-9]+)$")
		parsedBody := regexpObj.FindStringSubmatch(body)
		if parsedBody == nil {
			fmt.Printf("ParseError: %#v\n", body)
			return
		}

		parsedDate, err := utils.ParseDateStr(parsedBody[2])
		if err != nil {
			fmt.Printf("ParseDateError: %#v\n", parsedBody[2])
			return
		}

		repo := repository.NewDBRepository(db)

		task := repo.Add(parsedBody[1], parsedDate, username)
		fmt.Printf("added: %#v\n", task)
		client := utils.NewTwitterClient()
		client.Reply(fmt.Sprintf("%v (%v)\n", parsedBody[1], parsedBody[2]), statusId)
	},
}

func init() {
	CmdHandler.AddCommand(cmdAdd)
}
