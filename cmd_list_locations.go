package main

func ListLocations(c Client, _ []byte) error {
	return c.send(LocationsMessage{"locations", Locations})
}
