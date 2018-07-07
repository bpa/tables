package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/bpa/tables/data"
)

var siteUrl string
var Notifiers = make(map[string]data.Notifier)

type ImplEntry struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

func Initialize() {
}

func NotifyNewTable(table *data.Table, player *data.Player) {
	for name, notifications := range data.GetNotifications() {
		notifier := Notifiers[name]
		if notifier != nil {
			notifier.NotifyNewTable(table, player, &notifications)
		}
	}
}

func ConfigureNotifications(notifications map[string]ImplEntry) {
	for k := range notifications {
		var err error
		var note data.Notifier
		notifier := notifications[k]
		switch notifier.Type {
		case "http":
			note, err = NewNotifyHttp(notifier.Config)
		case "smtp":
			note, err = NewNotifySmtp(notifier.Config)
		default:
			err = errors.New(fmt.Sprintf("Unknown type '%s'", notifier.Type))
		}
		if err != nil {
			log.Fatal(fmt.Sprintf("Error in config.json/notifications/%s: %s", k, err.Error()))
		}
		Notifiers[k] = note
	}
}
