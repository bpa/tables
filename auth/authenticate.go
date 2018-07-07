package auth

import (
	"fmt"

	"github.com/bpa/tables/data"
)

type passwordLoginMessage struct {
	Username string
	Password string
}

type Authentication interface {
	Authenticate(data.Client, map[string]interface{}) (*data.Player, error)
}

var authImpl = make(map[string]Authentication)

func init() {
	authImpl["trusted"], _ = NewTrustedAuth()
}

func ConfigureAuthentication(authentication map[string]interface{}) {
	/*
		for k, auth := range authentication {
			var err error
			var impl Authentication
			method, _ := auth["Type"].(string)
			switch auth["Type"] {
			case "LDAPAuth":
				impl, err = NewLDAPAuth(auth.Config)
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
	*/
}

func Login(c data.Client, msg map[string]interface{}) (*data.Player, error) {
	method, ok := msg["method"].(string)
	if !ok {
		return nil, fmt.Errorf("Login missing field: `method`")
	}

	impl := authImpl[method]
	if impl == nil {
		return nil, fmt.Errorf("Unknown login method '%s'", method)
	}

	player, err := impl.Authenticate(c, msg)
	if err != nil {
		return nil, err
	}

	return player, nil
}
