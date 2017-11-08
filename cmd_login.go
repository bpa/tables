package main

import (
	"encoding/json"
	"errors"
	"github.com/satori/go.uuid"
)

type LoginMessage struct {
	Cmd    string `json:"cmd"`
	Player Player `json:"player"`
}

func Login(c *Client, msg []byte) error {
	var cmd LoginMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	if cmd.Player.FullName == "" {
		return errors.New("fullName is required")
	}

	cmd.Player.Id = uuid.NewV4().String()
	c.player = cmd.Player

	g, _ := json.Marshal(cmd)
	c.send <- g
	return nil
}
