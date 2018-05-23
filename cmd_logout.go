package main

func Logout(c Client, msg []byte) error {
	c.setPlayer(&Player{})

	c.send(CommandMessage{"logout"})
	return nil
}
