package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bpa/tables/data"
)

type NotifyHttp struct {
	Url     string `json:"url"`
	Message string `json:"message"`
}

func NewNotifyHttp(conf json.RawMessage) (NotifyHttp, error) {
	var nh NotifyHttp
	err := json.Unmarshal(conf, &nh)
	if err != nil {
		return nh, err
	}
	if nh.Url == "" {
		return nh, errors.New("Missing 'url'")
	}
	if nh.Message == "" {
		return nh, errors.New("Missing 'message'")
	}
	return nh, nil
}

func (nh NotifyHttp) NotifyNewTable(table *data.Table, author *data.Player, players *[]data.Player) {
	msg := fmt.Sprintf("%s created a table for %s at %s (%s)",
		author.FullName, table.Game.Name, table.Start.Local().Format(time.Kitchen), siteUrl)
	for _, p := range *players {
		nh.notify(&p, msg)
	}
}

func (nh NotifyHttp) notify(p *data.Player, message string) {
	http.Post(fmt.Sprintf(nh.Url, p.Id), "application/json", strings.NewReader(fmt.Sprintf(nh.Message, message)))
}
