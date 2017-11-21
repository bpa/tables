package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type NotifyHttp struct {
	Url     string
	Message string
}

func NewNotifyHttp(conf map[string]string) (NotifyHttp, error) {
	nh := NotifyHttp{conf["url"], conf["message"]}
	if nh.Url == "" {
		return nh, errors.New("Missing 'url'")
	}
	if nh.Message == "" {
		return nh, errors.New("Missing 'message'")
	}
	return nh, nil
}

func (nh NotifyHttp) notify(p *Player, message string) {
	http.Post(fmt.Sprintf(nh.Url, p.Id), "application/json", strings.NewReader(fmt.Sprintf(nh.Message, message)))
}
