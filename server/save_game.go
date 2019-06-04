package server

import (
	"errors"

	"github.com/bpa/tables/data"
	"github.com/mitchellh/mapstructure"
)

func init() {
	commands["save_game"] = saveGame
}

type saveGameMessage struct {
	Game data.Game
}

func saveGame(c data.Client, msg obj) error {
	var cmd saveGameMessage
	mapstructure.Decode(msg, &cmd)
	g := &cmd.Game

	if g.Name == "" {
		return errors.New("game.name is required")
	}

	if g.Min == 0 {
		return errors.New("game.min is required")
	}

	if g.Max == 0 {
		return errors.New("game.max is required")
	}

	if err := data.UpdateGame(g.ID, g.Name, g.Min, g.Max); err != nil {
		return err
	}
	return hub.Broadcast(gamesMessage{"games", data.GetGames()})
}
