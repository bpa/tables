package main

import (
	"encoding/json"
	"errors"
)

type TrustedAuth struct{}

func NewTrustedAuth() (TrustedAuth, error) {
	return TrustedAuth{}, nil
}

func (auth TrustedAuth) authenticate(c Client, msg []byte) (*Player, error) {
	var cmd passwordLoginMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return nil, err
	}

	if cmd.Username == "" {
		return nil, errors.New("username is required")
	}

	player := Player{FullName: cmd.Username, Id: c.host()}
	return &player, nil
}
