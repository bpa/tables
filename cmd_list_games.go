package main

func ListGames(c Client, _ []byte) error {
	return c.send(GamesMessage{"games", Games})
}
