package main

import (
	"encoding/json"
	"errors"
)

type TrustedAuth struct{}

func NewTrustedAuth() (TrustedAuth, error) {
	return TrustedAuth{}, nil
}

type noAuthLoginMessage struct {
	Cmd    string `json:"cmd"`
	Player Player `json:"player"`
}

func (auth TrustedAuth) authenticate(c *Client, msg []byte) (*Player, error) {
	var cmd noAuthLoginMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return nil, err
	}

	if cmd.Player.FullName == "" {
		return nil, errors.New("fullName is required")
	}

	cmd.Player.Id = c.remoteHost
	return &c.player, nil
}
