package tables

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Commands map[string]func([]byte) ([]byte, error)

var commands Commands = make(Commands)

func get(f func() interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		out, err := json.Marshal(f())
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			w.Write(out)
		}
	}
}

func post(f func(interface{}) interface{}, t func(*json.Decoder) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		val, err := t(decoder)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		out, err := json.Marshal(f(val))
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			w.Write(out)
		}
	}
}

func Listen(addr string) {
	Tables = append(Tables, Table{Name: "test"})
	r := mux.NewRouter()
	r.Methods("GET").Path("/table").HandlerFunc(get(GetTables))
	r.Methods("POST").Path("/table").HandlerFunc(post(AddTable, DecodeTable))
	http.Handle("/", r)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
