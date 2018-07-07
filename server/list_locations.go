package server

import "github.com/bpa/tables/data"

func init() {
	commands["list_locations"] = listLocations
}

func listLocations(c data.Client, _ obj) error {
	c.Send(locationsMessage{"locations", data.GetLocations()})
	return nil
}
