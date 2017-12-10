package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/ldap.v2"
)

type LDAPAuth struct {
	Addr           string `json:"addr"`
	BaseDN         string `json:"baseDN"`
	Filter         string `json:"filter"`
	Network        string `json:"network"`
	UsernameFormat string `json:"usernameFormat"`
	Email          string `json:"email"`
	FirstName      string `json:"firstName"`
	FullName       string `json:"FullName"`
}

func getOrDefault(m map[string]string, k, d string) string {
	v, ok := m[k]
	if ok {
		return v
	}
	return d
}

func NewLDAPAuth(config *json.RawMessage) (LDAPAuth, error) {
	var auth LDAPAuth
	err := json.Unmarshal(*config, &auth)
	if err != nil {
		return auth, err
	}
	if len(auth.Network) == 0 {
		auth.Network = "tcp"
	}
	if len(auth.Addr) == 0 {
		auth.Addr = "ldap:389"
	}
	if len(auth.UsernameFormat) == 0 {
		auth.UsernameFormat = "%s"
	}
	if len(auth.Filter) == 0 {
		auth.Filter = "(sAMAccountName=%s)"
	}
	if len(auth.BaseDN) == 0 {
		return auth, errors.New("Missing baseDN")
	}
	return auth, nil
}

func (auth LDAPAuth) authenticate(client Client, packet []byte) (*Player, error) {
	var msg passwordLoginMessage
	err := json.Unmarshal(packet, &msg)
	if err != nil {
		return nil, err
	}

	l, err := ldap.Dial(auth.Network, auth.Addr)
	if err != nil {
		return nil, err
	}
	defer l.Close()

	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, err
	}

	err = l.Bind(fmt.Sprintf(auth.UsernameFormat, msg.Username), msg.Password)
	if err != nil {
		return nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		auth.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(auth.Filter, msg.Username),
		[]string{
			auth.Email,
			auth.FirstName,
			auth.FullName},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	if len(sr.Entries) != 1 {
		return nil, errors.New("User does not exist")
	}

	p := sr.Entries[0]
	return &Player{
			p.GetAttributeValue(auth.FirstName),
			p.GetAttributeValue(auth.FullName),
			p.GetAttributeValue(auth.Email),
			p.GetAttributeValue(auth.Email)},
		nil
}
