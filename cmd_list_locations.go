package main

import "encoding/json"

func ListLocations(c *Client, _ []byte) error {
	g, _ := json.Marshal(LocationsMessage{"locations", Locations})
	c.send <- g
	return nil
}
