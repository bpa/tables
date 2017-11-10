package main

import (
	"encoding/json"
	"errors"
	"strings"
)

type LoginMessage struct {
	Cmd    string `json:"cmd"`
	Player Player `json:"player"`
}

func NoAuthLogin(c *Client, msg []byte) error {
	var cmd LoginMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	if cmd.Player.FullName == "" {
		return errors.New("fullName is required")
	}

	addr := c.conn.RemoteAddr().String()
	i := strings.LastIndexByte(addr, ':')
	if i != -1 {
		addr = addr[:i]
	}
	cmd.Player.Id = addr
	c.player = cmd.Player

	g, _ := json.Marshal(cmd)
	c.send <- g
	return nil
}
