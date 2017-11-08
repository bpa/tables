package main

import "time"

func startCleaner() {
	ticker := time.NewTicker(time.Minute * 15)
	go func() {
		for range ticker.C {
			RemoveExpiredTables()
		}
	}()
}
