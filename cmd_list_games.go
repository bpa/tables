package main

import "encoding/json"

func ListGames(c *Client, _ []byte) error {
	g, _ := json.Marshal(GamesMessage{"games", Games})
	c.send <- g
	return nil
}
