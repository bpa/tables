package notify

import (
	"errors"
	"fmt"
	"log"

	"github.com/bpa/tables/data"
	"github.com/spf13/viper"
)

var siteUrl string
var Notifiers = make(map[string]data.Notifier)

func Initialize() {
	siteUrl = viper.GetString("SiteUrl")
	n := viper.Get("notifications")
	conf := n.(map[string]interface{})
	ConfigureNotifications(conf)
}

func getString(vc map[string]interface{}, key string) string {
	value, ok := vc[key].(string)
	if ok {
		return value
	}
	return ""
}

func NotifyNewTable(table *data.Table, player *data.Player) {
	for name, notifications := range data.GetNotifications() {
		notifier := Notifiers[name]
		if notifier != nil {
			notifier.NotifyNewTable(table, player, &notifications)
		}
	}
}

func ConfigureNotifications(notifications map[string]interface{}) {
	for k := range notifications {
		var err error
		var note data.Notifier
		n := notifications[k]
		notifier := n.(map[string]interface{})
		notifierType := notifier["type"]
		notifierConfig := notifier["config"].(map[string]interface{})
		switch notifierType {
		case "http":
			note, err = NewNotifyHttp(notifierConfig)
		case "smtp":
			note, err = NewNotifySmtp(notifierConfig)
		default:
			err = errors.New(fmt.Sprintf("Unknown type '%s'", notifierType))
		}
		if err != nil {
			log.Fatal(fmt.Sprintf("Error in config.json/notifications/%s: %s", k, err.Error()))
		}
		Notifiers[k] = note
	}
}
