package main

import "encoding/json"

type LocationsMessage struct {
	Cmd       string   `json:"cmd"`
	Locations []string `json:"locations"`
}

type GamesMessage struct {
	Cmd   string          `json:"cmd"`
	Games map[string]Game `json:"games"`
}

func NewGame(c *Client, _ []byte) error {
	l, _ := json.Marshal(LocationsMessage{"locations", Locations})
	c.send <- l

	g, _ := json.Marshal(GamesMessage{"games", Games})
	c.send <- g
	return nil
}
