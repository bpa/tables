package main

import (
	"encoding/json"
	"errors"
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

	c.player = cmd.Player

	g, _ := json.Marshal(cmd)
	c.hub.broadcast <- g
	return nil
}
