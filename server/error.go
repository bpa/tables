package server

import (
	"fmt"

	"github.com/bpa/tables/data"
)

func init() {
	commands["error"] = printError
}

func printError(c data.Client, message obj) error {
	if err := message["error"]; err != "" {
		fmt.Println(err)
	}
	return nil
}
