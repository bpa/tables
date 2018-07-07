package server

import "github.com/bpa/tables/data"

func init() {
	commands["logout"] = logout
}

func logout(c data.Client, msg obj) error {
	c.SetPlayer(&data.Player{})

	c.Send(CommandMessage{"logout"})
	return nil
}
