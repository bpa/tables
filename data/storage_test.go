package data

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"
)

func forImpl(f func(string, Storage)) {
	for _, i := range []struct {
		name    string
		storage Storage
	}{
		{"MemoryStorage", NewMemoryStorage()},
	} {
		f(i.name, i.storage)
	}
}

type byName []Game

func (g byName) Len() int           { return len(g) }
func (g byName) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g byName) Less(i, j int) bool { return g[i].Name < g[j].Name }

func TestGame(t *testing.T) {
	forImpl(func(name string, storage Storage) {
		if len(*storage.GetGames()) != 0 {
			t.Fatalf("%s: not initialized, len(games) != 0", name)
			return
		}
		storage.CreateGame("1", 1, 3)

		if len(*storage.GetGames()) != 1 {
			t.Fatalf("%s: Game not persisted", name)
			return
		}

		storage.CreateGame("2", 2, 4)
		storage.CreateGame("3", 3, 5)
		list := storage.GetGames()
		if len(*list) != 3 {
			t.Fatalf("%s: Wanted %d games, got %d", name, 3, len(*list))
			return
		}

		sort.Sort(byName(*list))
		for i, g := range *list {
			if g.Min != i+1 || g.Max != i+3 {
				t.Errorf("%s: %d-%d != %d-%d for %d", name, g.Min, g.Max, i+1, i+3, i+1)
			}
			storage.UpdateGame(g.Id, g.Name+"+2", i+2, i+4)
		}
		list = storage.GetGames()
		sort.Sort(byName(*list))
		if len(*list) != 3 {
			t.Errorf("After update, have %d games, wanted 3", len(*list))
		}
		for i, g := range *list {
			if g.Name != fmt.Sprintf("%d+2", i+1) {
				t.Errorf("Name not updated")
			}
			if g.Min != i+2 {
				t.Errorf("Min not updated")
			}
			if g.Max != i+4 {
				t.Errorf("Max not updated")
			}
		}
		storage.DeleteGame((*list)[1].Id)
		list = storage.GetGames()
		sort.Sort(byName(*list))
		if len(*list) != 2 {
			t.Fatalf("After delete, have %d games, wanted 2", len(*list))
		}
		if (*list)[0].Name != "1+2" {
			t.Errorf("game 1 deleted by mistake")
		}
		if (*list)[1].Name != "3+2" {
			t.Errorf("game 3 deleted by mistake")
		}
		storage.DeleteGame((*list)[0].Id)
		storage.DeleteGame((*list)[1].Id)
		if len(*storage.GetGames()) != 0 {
			t.Errorf("Deletion failed")
		}
	})
}

func checkLocations(t *testing.T, storage Storage, name string, exp *[]string) {
	t.Helper()
	list := storage.GetLocations()
	sort.Strings(*list)
	if !reflect.DeepEqual(list, exp) {
		t.Errorf("%s: Locations don't match.  Wanted %v, got %v", name, exp, list)
	}
}

func TestLocation(t *testing.T) {
	forImpl(func(name string, storage Storage) {
		checkLocations(t, storage, name, &[]string{})
		storage.CreateLocation("a")
		storage.CreateLocation("c")
		storage.CreateLocation("b")

		checkLocations(t, storage, name, &[]string{"a", "b", "c"})
		storage.DeleteLocation("c")
		checkLocations(t, storage, name, &[]string{"a", "b"})
		storage.UpdateLocation("a", "z")
		checkLocations(t, storage, name, &[]string{"b", "z"})
		storage.DeleteLocation("a")
		checkLocations(t, storage, name, &[]string{"b", "z"})
		storage.DeleteLocation("b")
		storage.DeleteLocation("z")
		checkLocations(t, storage, name, &[]string{})
	})
}

func assertEq(t *testing.T, name, data string, exp, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(exp, actual) {
		t.Errorf("%s: %s don't match.\nWanted: %v\n   got: %v", name, data, exp, actual)
	}
}

func TestPersistence(t *testing.T) {
	defer os.Remove("state.json")
	now := time.Now().Round(time.Nanosecond)
	creator := Player{FirstName: "tester", Id: "me"}

	ids := make(map[string]string)
	tableIds := make(map[string]string)
	forImpl(func(name string, storage Storage) {
		storage.CreateLocation(name)
		storage.CreateGame("play", 1, 5)
		ids[name] = (*storage.GetGames())[0].Id
		storage.CreateTable(ids[name], name, now, &creator)
		tableIds[name] = (*storage.GetTables())[0].Id
		storage.SetPlayerNotifications(&creator, []string{"test"})
	})

	forImpl(func(name string, storage Storage) {
		checkLocations(t, storage, name, &[]string{name})

		testGame := Game{"play", 1, 5, ids[name]}
		expGames := []Game{testGame}
		assertEq(t, name, "Games", &expGames, storage.GetGames())

		expTables := []Table{Table{
			Game:     &testGame,
			Players:  []*Player{&creator},
			Location: name,
			Start:    now,
			Id:       tableIds[name],
		}}
		assertEq(t, name, "Tables", &expTables, storage.GetTables())

		expNotifications := map[string][]Player{
			"test": []Player{creator},
		}
		assertEq(t, name, "Notifications", expNotifications, storage.GetNotifications())
	})
}

/*
	CreateTable(game, location string, start time.Time, p *Player) (*Table, error)
	DeleteExpiredTables() bool
	DeleteTable(string) error
	GetNotifications() map[string][]Player
	GetTables() *[]Table
	JoinTable(*Player, string) error
	LeaveTable(*Player, string) error
*/
