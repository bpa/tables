package auth

import (
	"errors"

	"github.com/bpa/tables/data"
)

type TrustedAuth struct{}

func NewTrustedAuth() (TrustedAuth, error) {
	return TrustedAuth{}, nil
}

func (auth TrustedAuth) Authenticate(c data.Client, msg map[string]interface{}) (*data.Player, error) {
	username, _ := msg["username"].(string)
	if username == "" {
		return nil, errors.New("username is required")
	}

	player := data.Player{FullName: username, Id: c.Host()}
	return &player, nil
}
