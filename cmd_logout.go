package main

import (
	"encoding/json"
)

func Logout(c *Client, msg []byte) error {
	c.player = Player{}

	g, _ := json.Marshal(CommandMessage{"logout"})
	c.send <- g
	return nil
}
