package server

import (
	"errors"

	"github.com/bpa/tables/data"
)

func init() {
	commands["delete_game"] = deleteGame
}

func deleteGame(c data.Client, msg obj) error {
	game, _ := msg["game"].(string)
	if game == "" {
		return errors.New("Game is required")
	}
	if err := data.DeleteGame(game); err != nil {
		return err
	}
	return hub.Broadcast(gamesMessage{"games", data.GetGames()})
}
