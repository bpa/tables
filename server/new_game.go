package server

import "github.com/bpa/tables/data"

func init() {
	commands["new_game"] = newGame
}

type locationsMessage struct {
	Cmd       string    `json:"cmd"`
	Locations *[]string `json:"locations"`
}

type gamesMessage struct {
	Cmd   string       `json:"cmd"`
	Games *[]data.Game `json:"games"`
}

type tablesMessage struct {
	Cmd    string        `json:"cmd"`
	Tables *[]data.Table `json:"tables"`
}

func newGame(c data.Client, _ obj) error {
	c.Send(locationsMessage{"locations", data.GetLocations()})
	c.Send(gamesMessage{"games", data.GetGames()})
	return nil
}
