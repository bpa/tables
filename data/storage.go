package data

import "time"

type Storage interface {
	CreateGame(name string, min, max int) error
	CreateLocation(string) error
	CreateTable(game, location string, start time.Time, p *Player) (*Table, error)
	DeleteExpiredTables() bool
	DeleteGame(string) error
	DeleteLocation(string) error
	DeletePlayerNotifications(*Player)
	DeleteTable(string) error
	GetGames() *[]Game
	GetLocations() *[]string
	GetNotifications() map[string][]Player
	GetPlayerNotifications(*Player) []string
	GetTables() *[]Table
	JoinTable(*Player, string) error
	LeaveTable(*Player, string) error
	SetPlayerNotifications(p *Player, methods []string)
	UpdateGame(id, name string, min, max int) error
	UpdateLocation(from, to string) error
}
