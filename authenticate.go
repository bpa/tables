package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type authenticationMessage struct {
	Type string `json:"type"`
}

type passwordLoginMessage struct {
	Cmd      string `json:"cmd"`
	Type     string `json:"Type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type authenticatedResponse struct {
	Cmd    string `json:"cmd"`
	Player Player `json:"player"`
}

type Authentication interface {
	authenticate(Client, []byte) (*Player, error)
}

var authImpl = make(map[string]Authentication)

func ConfigureAuthentication(authentication map[string]ImplEntry) {
	for k := range authentication {
		var err error
		auth := authentication[k]
		var impl Authentication
		switch auth.Type {
		case "LDAPAuth":
			impl, err = NewLDAPAuth(auth.config)
		case "TrustedAuth":
			impl, err = NewTrustedAuth()
		default:
			log.Fatal("Unknown authentication implementation '%s'", auth.Type)
		}
		if err != nil {
			log.Fatal("Can't create '%s' of type %s: %s", k, auth.Type, err.Error())
		}
		authImpl[k] = impl
	}
}

func Login(c Client, msg []byte) error {
	var auth authenticationMessage
	err := json.Unmarshal(msg, &auth)
	if err != nil {
		return err
	}

	impl := authImpl[auth.Type]
	if impl == nil {
		return errors.New(fmt.Sprintf("Unknown login type '%s'", auth.Type))
	}

	player, err := impl.authenticate(c, msg)
	if err != nil {
		return err
	}
	res := authenticatedResponse{"login", *player}
	c.setPlayer(player)

	g, _ := json.Marshal(res)
	c.send(&g)
	return nil
}
