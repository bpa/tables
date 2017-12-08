package main

type PasswordLoginMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authentication interface {
	authenticate(client *Client, packet []byte) (*Player, error)
}

//	c.player = cmd.Player
//
//	g, _ := json.Marshal(cmd)
//	c.send <- g
//	return nil
