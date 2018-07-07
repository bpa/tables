package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bpa/tables/data"
)

var hub Hub

func Listen(port int) {
	wsHub := newWSHub()
	hub = wsHub
	go wsHub.run()

	ticker := time.NewTicker(time.Minute * 15)
	go func() {
		for range ticker.C {
			if data.DeleteExpiredTables() {
				hub.Broadcast(data.GetTables())
			}
		}
	}()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(wsHub, w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("dist")))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
