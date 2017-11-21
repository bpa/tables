package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/ldap.v2"
)

type LDAPConfig struct {
	addr           string
	baseDN         string
	filter         string
	network        string
	usernameFormat string
	email          string
	firstName      string
	fullName       string
}

var ldapConfig LDAPConfig

func getOrDefault(m map[string]string, k, d string) string {
	v, ok := m[k]
	if ok {
		return v
	}
	return d
}

func LDAPInit(conf map[string]string) {
	ldapConfig.network = getOrDefault(conf, "network", "tcp")
	ldapConfig.addr = getOrDefault(conf, "addr", "ldap:389")
	ldapConfig.usernameFormat = getOrDefault(conf, "userFormat", "%s")
	ldapConfig.filter = getOrDefault(conf, "filter", "(sAMAccountName=%s)")
	ldapConfig.baseDN = conf["baseDN"]
	if ldapConfig.baseDN == "" {
		log.Fatal("Missing authentication/baseDN in config.json")
	}
}

func LDAPLogin(*Client, []byte) error {
	return nil
}

func authenticateLDAP(username, password string) (*Player, error) {
	l, err := ldap.Dial(ldapConfig.network, ldapConfig.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, err
	}

	err = l.Bind(fmt.Sprintf(ldapConfig.usernameFormat, username), password)
	if err != nil {
		return nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		ldapConfig.baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(ldapConfig.filter, username),
		[]string{
			ldapConfig.email,
			ldapConfig.firstName,
			ldapConfig.fullName},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	p := sr.Entries[0]
	return &Player{
			p.GetAttributeValue(ldapConfig.firstName),
			p.GetAttributeValue(ldapConfig.fullName),
			p.GetAttributeValue(ldapConfig.email),
			p.GetAttributeValue(ldapConfig.email)},
		nil
}
