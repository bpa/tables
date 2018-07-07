package server

import (
	"os"
	"testing"

	"github.com/bpa/tables/test"
)

type TestHub struct {
	sent []interface{}
}

var testHub = &TestHub{make([]interface{}, 0, 3)}

func (h *TestHub) Broadcast(v interface{}) error {
	h.sent = append(h.sent, v)
	return nil
}

func TestMain(m *testing.M) {
	hub = testHub
	os.Exit(m.Run())
}

func TestSend(t *testing.T) {
	c := test.NewClient(HandleMessage)
	c.Message(CommandMessage{"not-a-thing"})
	c.AssertError(t, "Invalid command: 'not-a-thing'")
}

func TestBroadcast(t *testing.T) {
	hub.Broadcast(CommandMessage{"testing"})
	if len(testHub.sent) != 1 {
		t.Errorf("Expected 1 broadcasted message, found %d", len(testHub.sent))
	}
	c, _ := testHub.sent[0].(CommandMessage)
	if c.Cmd != "testing" {
		t.Errorf("Expected command `testing`, found `%s`", c.Cmd)
	}
}
