package main

import (
	"encoding/json"
	"fmt"
)

type CommandMessage struct {
	Cmd string `json:"cmd"`
}

type Command func(*Client, []byte) error

type ErrorMessage struct {
	Cmd     string `json:"cmd"`
	Message string `json:"message"`
}

func PrintError(c *Client, msg []byte) error {
	var err ErrorMessage
	json.Unmarshal(msg, &err)
	fmt.Println(err.Message)
	return nil
}

var commands = map[string]Command{
	"create_table": CreateTable,
	"error":        PrintError,
	"join_table":   JoinTable,
	"leave_table":  LeaveTable,
	"list_games":   ListGames,
	"login":        Login,
	"logout":       Logout,
	"new_game":     NewGame,
}

func handleMessage(c *Client, message []byte) {
	var cmd CommandMessage
	json.Unmarshal(message, &cmd)
	f := commands[cmd.Cmd]
	if f == nil {
		fmt.Println(cmd.Cmd)
		err, _ := json.Marshal(ErrorMessage{
			"error", fmt.Sprintf("Invalid command: '%s'", cmd.Cmd)})
		c.send <- err
	} else {
		err := f(c, message)
		if err != nil {
			e, _ := json.Marshal(ErrorMessage{"error", err.Error()})
			c.send <- e
		}
	}
}
