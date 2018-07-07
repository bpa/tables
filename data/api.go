package data

import "time"

var storage Storage

const EXPIRE_TIME = time.Minute * 45

func Initialize() {
	storage = NewMemoryStorage()
}

func CreateGame(name string, min, max int) error {
	return storage.CreateGame(name, min, max)
}

func CreateLocation(location string) error {
	return storage.CreateLocation(location)
}

func CreateTable(game, location string, start time.Time, p *Player) (*Table, error) {
	return storage.CreateTable(game, location, start, p)
}

func DeleteExpiredTables() bool {
	return storage.DeleteExpiredTables()
}

func DeleteGame(game string) error {
	return storage.DeleteGame(game)
}

func DeleteLocation(location string) error {
	return storage.DeleteLocation(location)
}

func DeletePlayerNotifications(p *Player) {
	storage.DeletePlayerNotifications(p)
}

func DeleteTable(table string) error {
	return storage.DeleteTable(table)
}

func GetGames() *[]Game {
	return storage.GetGames()
}

func GetLocations() *[]string {
	return storage.GetLocations()
}

func GetNotifications() map[string][]Player {
	return storage.GetNotifications()
}

func GetPlayerNotifications(p *Player) []string {
	return storage.GetPlayerNotifications(p)
}

func GetTables() *[]Table {
	return storage.GetTables()
}

func JoinTable(p *Player, table string) error {
	return storage.JoinTable(p, table)
}

func LeaveTable(p *Player, table string) error {
	return storage.LeaveTable(p, table)
}

func SetPlayerNotifications(p *Player, methods []string) {
	storage.SetPlayerNotifications(p, methods)
}

func UpdateGame(id, name string, min, max int) error {
	return storage.UpdateGame(id, name, min, max)
}

func UpdateLocation(from, to string) error {
	return storage.UpdateLocation(from, to)
}
