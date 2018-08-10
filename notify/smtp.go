package notify

import (
	"errors"
	"fmt"
	"net/smtp"
	"time"

	"github.com/bpa/tables/data"
)

type NotifySmtp struct {
	Host  string
	Email string
}

func NewNotifySmtp(conf map[string]interface{}) (NotifySmtp, error) {
	var ns NotifySmtp
	ns.Host = getString(conf, "host")
	if ns.Host == "" {
		return ns, errors.New("Missing 'host'")
	}
	ns.Email = getString(conf, "email")
	if ns.Email == "" {
		return ns, errors.New("Missing 'email'")
	}
	return ns, nil
}

func (ns NotifySmtp) NotifyNewTable(table *data.Table, author *data.Player, players *[]data.Player) {
	msg := fmt.Sprintf("From: %s\nSubject: New Table\n\n%s created a table for %s at %s (%s)",
		ns.Email, author.FullName, table.Game.Name, table.Start.Local().Format(time.Kitchen), siteUrl)
	c, err := smtp.Dial(ns.Host)
	if err != nil {
		return
	}
	for _, p := range *players {
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
