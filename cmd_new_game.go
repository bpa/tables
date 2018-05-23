package main

type LocationsMessage struct {
	Cmd       string   `json:"cmd"`
	Locations []string `json:"locations"`
}

type GamesMessage struct {
	Cmd   string          `json:"cmd"`
	Games map[string]Game `json:"games"`
}

func NewGame(c Client, _ []byte) error {
	c.send(LocationsMessage{"locations", Locations})
	c.send(GamesMessage{"games", Games})
	return nil
}
