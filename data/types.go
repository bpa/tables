package data

import "time"

type Player struct {
	FirstName string `json:"firstName"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	Id        string `json:"id"`
}

type Game struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Id   string `json:"id"`
}

type Table struct {
	Game     *Game     `json:"game"`
	Players  []*Player `json:"players"`
	Location string    `json:"location"`
	Start    time.Time `json:"start"`
	Id       string    `json:"id"`
}
