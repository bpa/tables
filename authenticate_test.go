package main

import "testing"

type insecureAuth struct{}

func (auth insecureAuth) authenticate(c Client, msg []byte) (*Player, error) {
	return nil, nil
}

func TestAuth(t *testing.T) {
	authImpl["insecure"] = insecureAuth{}
	c := newTestClient()
	c.message(CommandMessage{"login"})
	err := assertError(&c, "Login missing type")
	if err != nil {
		t.Error(err)
	}
	t.Error(c.sent)
}
