package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type memoryFile struct {
	Games         []Game
	Locations     []string
	Tables        []Table
	Notifications map[string][]Player
}

type MemoryStorage struct {
	games         map[string]Game
	locations     map[string]bool
	notifications map[string]map[string]*Player
	playing       map[string]map[string]bool
	tables        map[string]*Table
}

func NewMemoryStorage() Storage {
	m := &MemoryStorage{
		games:         make(map[string]Game),
		locations:     make(map[string]bool),
		notifications: make(map[string]map[string]*Player),
		playing:       make(map[string]map[string]bool),
		tables:        make(map[string]*Table),
	}
	m.readState()
	return m
}

func (m *MemoryStorage) CreateGame(name string, min, max int) error {
	return m.UpdateGame("", name, min, max)
}

func (m *MemoryStorage) CreateLocation(location string) error {
	if _, exists := m.locations[location]; exists {
		return errors.New("Location already exists")
	}

	m.locations[location] = true

	m.saveState()
	return nil
}

func (m *MemoryStorage) CreateTable(game, location string, start time.Time, creator *Player) (*Table, error) {
	gameData := m.games[game]
	if gameData.Name == "" {
		return &Table{}, fmt.Errorf(fmt.Sprintf("No game %s available", game))
	}

	players := make([]*Player, 0, gameData.Max)
	players = append(players, creator)

	table := m.AddNewTable(&gameData, location, start, players)
	m.playing[table.ID] = make(map[string]bool)
	m.playing[table.ID][creator.ID] = true
	return &table, nil
}

func (m *MemoryStorage) DeleteGame(game string) error {
	_, ok := m.games[game]
	if ok {
		delete(m.games, game)
		m.saveState()
		return nil
	}
	return errors.New("Game does not exist")
}

func (m *MemoryStorage) DeleteLocation(location string) error {
	if m.locations[location] {
		delete(m.locations, location)
		m.saveState()
		return nil
	}
	return errors.New("Location does not exist")
}

func (m *MemoryStorage) DeletePlayerNotifications(p *Player) {
	for _, v := range m.notifications {
		delete(v, p.ID)
	}
}

func (m *MemoryStorage) DeleteTable(table string) error {
	if _, ok := m.tables[table]; ok {
		delete(m.tables, table)
		m.saveState()
		return nil
	}
	return errors.New("Table does not exist")
}

func (m *MemoryStorage) UpdateLocation(from, to string) error {
	if m.locations[from] {
		delete(m.locations, from)
		m.locations[to] = true
		m.saveState()
		return nil
	}
	return errors.New("Location does not exist")
}

func (m *MemoryStorage) JoinTable(player *Player, id string) error {
	table, ok := m.tables[id]
	if !ok {
		return errors.New("Unknown table")
	}

	if m.playing[id][player.ID] {
		return errors.New("Player already at table")
	}

	table.Players = append(table.Players, player)
	m.playing[id][player.ID] = true
	m.saveState()
	return nil
}

func (m *MemoryStorage) LeaveTable(player *Player, id string) error {
	table, ok := m.tables[id]
	if !ok {
		return errors.New("Unknown table")
	}

	if !m.playing[id][player.ID] {
		return errors.New("Player isn't at table")
	}

	delete(m.playing[id], player.ID)
	for i := range table.Players {
		if table.Players[i].ID == player.ID {
			table.Players = append(table.Players[:i], table.Players[i+1:]...)
			m.saveState()
			return nil
		}
	}

	return nil
}

func (m *MemoryStorage) UpdateGame(id, name string, min, max int) error {
	_, ok := m.games[id]
	if !ok {
		id = uuid.Must(uuid.NewRandom()).String()
	}

	m.games[id] = Game{name, min, max, id}
	m.saveState()
	return nil
}

func (m *MemoryStorage) readState() {
	f, err := os.Open("state.json")
	if err == nil {
		defer f.Close()

		var mem memoryFile
		dec := json.NewDecoder(f)
		err = dec.Decode(&mem)
		if err == nil {
			for _, t := range mem.Tables {
				table := new(Table)
				*table = t
				m.tables[t.ID] = table
				playing := make(map[string]bool)
				for _, p := range table.Players {
					playing[p.ID] = true
				}
				m.playing[t.ID] = playing
			}

			for _, g := range mem.Games {
				m.games[g.ID] = g
			}

			for _, l := range mem.Locations {
				m.locations[l] = true
			}

			for name, data := range mem.Notifications {
				players := make(map[string]*Player)
				for _, p := range data {
					player := new(Player)
					*player = p
					players[p.ID] = player
				}
				m.notifications[name] = players
			}
		} else {
			fmt.Printf("Error reading state.json: %s\n", err)
		}
	}
}

func (m *MemoryStorage) saveState() {
	f, err := os.OpenFile("state.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err == nil {
		defer f.Close()

		enc := json.NewEncoder(f)
		err = enc.Encode(memoryFile{
			Games:         *m.GetGames(),
			Locations:     *m.GetLocations(),
			Tables:        *m.GetTables(),
			Notifications: m.GetNotifications(),
		})
		if err != nil {
			fmt.Printf("Error writing state.json: %s\n", err)
		}
	}
}

func DecodeTable(d *json.Decoder) (interface{}, error) {
	var t Table
	err := d.Decode(&t)
	return t, err
}

func (m *MemoryStorage) GetGames() *[]Game {
	games := make([]Game, 0, len(m.games))
	for _, v := range m.games {
		games = append(games, v)
	}
	return &games
}

func (m *MemoryStorage) GetLocations() *[]string {
	locations := make([]string, 0, len(m.locations))
	for l := range m.locations {
		locations = append(locations, l)
	}
	return &locations
}

func (m *MemoryStorage) GetNotifications() map[string][]Player {
	notifications := make(map[string][]Player)
	for k, v := range m.notifications {
		players := make([]Player, 0, len(v))
		for _, p := range v {
			players = append(players, *p)
		}
		notifications[k] = players
	}
	return notifications
}

//TODO: implement
func (m *MemoryStorage) GetPlayerNotifications(p *Player) []string {
	return nil
}

func (m *MemoryStorage) SetPlayerNotifications(p *Player, n []string) {
	for _, method := range n {
		notifications, ok := m.notifications[method]
		if !ok {
			notifications = make(map[string]*Player)
			m.notifications[method] = notifications
		}
		notifications[p.ID] = p
	}
	m.saveState()
}

func (m *MemoryStorage) GetTables() *[]Table {
	tables := make([]Table, 0, len(m.tables))
	for _, v := range m.tables {
		tables = append(tables, *v)
	}
	return &tables
}

func (m *MemoryStorage) AddNewTable(game *Game, loc string, start time.Time, players []*Player) Table {
	id := uuid.Must(uuid.NewRandom()).String()
	table := Table{
		Game:     game,
		Players:  players,
		Location: loc,
		Start:    start,
		ID:       id,
	}
	m.tables[id] = &table
	m.playing[id] = make(map[string]bool)
	m.saveState()
	return table
}

func (m *MemoryStorage) DeleteExpiredTables() bool {
	removed := false
	now := time.Now()
	for _, t := range m.tables {
		if t.Start.Add(EXPIRE_TIME).Before(now) {
			delete(m.tables, t.ID)
			delete(m.playing, t.ID)
			removed = true
		}
	}
	if removed {
		m.saveState()
	}
	return removed
}
