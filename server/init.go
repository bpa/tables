package server

import (
	"encoding/json"
	"fmt"

	"github.com/bpa/tables/data"
)

type obj map[string]interface{}

var commands = make(map[string]func(data.Client, obj) error)

type CommandMessage struct {
	Cmd string `json:"cmd"`
}

func HandleMessage(c data.Client, message []byte) {
	var msg obj
	if err := json.Unmarshal(message, &msg); err != nil {
		c.Error("Invalid message")
		return
	}

	cmd, _ := msg["cmd"].(string)
	f := commands[cmd]
	if f == nil {
		fmt.Println(cmd)
		c.Error("Invalid command: '%s'", cmd)
		return
	}

	if err := f(c, msg); err != nil {
		c.Error(err.Error())
	}
}

func handleConnect(c data.Client) {
	c.Send(tablesMessage{"tables", data.GetTables()})
}

func authenticated(f func(data.Client, obj) error) func(data.Client, obj) error {
	return func(c data.Client, msg obj) error {
		if c.GetPlayer() == nil {
			cmd, _ := msg["cmd"].(string)
			return fmt.Errorf("Authentication required for `%s`", cmd)
		}
		return f(c, msg)
	}
}
