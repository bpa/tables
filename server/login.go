package server

import (
	"github.com/bpa/tables/auth"
	"github.com/bpa/tables/data"
)

type authenticatedResponse struct {
	Cmd    string       `json:"cmd"`
	Player *data.Player `json:"player"`
}

func init() {
	commands["login"] = login
}

func login(c data.Client, msg obj) error {
	player, err := auth.Login(c, msg)
	if err != nil {
		return err
	}
	c.SetPlayer(player)
	c.Send(authenticatedResponse{"login", player})
	return nil
}
