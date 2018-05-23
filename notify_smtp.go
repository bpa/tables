package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/smtp"
	"time"
)

type NotifySmtp struct {
	Host  string
	Email string
}

func NewNotifySmtp(conf json.RawMessage) (NotifySmtp, error) {
	var ns NotifySmtp
	err := json.Unmarshal(conf, &ns)
	if err != nil {
		return ns, err
	}
	if ns.Host == "" {
		return ns, errors.New("Missing 'host'")
	}
	if ns.Email == "" {
		return ns, errors.New("Missing 'email'")
	}
	return ns, nil
}

func (ns NotifySmtp) notifyNewTable(table *Table, author *Player, players []Player) {
	msg := fmt.Sprintf("Subject: New Table\n\n%s created a table for %s at %s (%s)",
		author.FullName, table.Game.Name, table.Start.Local().Format(time.Kitchen), SiteUrl)
	c, err := smtp.Dial(ns.Host)
	if err != nil {
		return
	}
	for _, p := range players {
		if err := c.Mail(ns.Email); err != nil {
			continue
		}
		if err := c.Rcpt(p.Email); err != nil {
			continue
		}
		wc, err := c.Data()
		if err != nil {
			continue
		}
		_, err = fmt.Fprintf(wc, msg)
		if err != nil {
			continue
		}
		wc.Close()
	}
}
