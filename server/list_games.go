package server

import "github.com/bpa/tables/data"

func init() {
	commands["list_games"] = listGames
}

func listGames(c data.Client, _ obj) error {
	c.Send(gamesMessage{"games", data.GetGames()})
	return nil
}
