package data

import "time"

type Player struct {
	FirstName string `json:"firstName"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	ID        string `json:"id"`
}

type Game struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	ID   string `json:"id"`
}

type Table struct {
	Game     *Game     `json:"game"`
	Players  []*Player `json:"players"`
	Location string    `json:"location"`
	Start    time.Time `json:"start"`
	ID       string    `json:"id"`
}
