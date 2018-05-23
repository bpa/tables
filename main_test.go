package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type testClient struct {
	sent   []interface{}
	player *Player
}

func (tc *testClient) send(v interface{}) error {
	tc.sent = append(tc.sent, v)
	return nil
}

func (tc *testClient) message(v interface{}) error {
	g, _ := json.Marshal(v)
	handleMessage(tc, g)
	return nil
}

func (tc *testClient) host() string {
	return "localhost"
}

func (tc *testClient) setPlayer(player *Player) {
	tc.player = player
}

func newTestClient() *testClient {
	return &testClient{sent: make([]interface{}, 0, 3)}
}

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

func assertError(c *testClient, msg string) error {
	m, ok := c.sent[0].(ErrorMessage)
	if !ok {
		err := fmt.Sprintf("Expecting CommandMessage, found `%s`", reflect.TypeOf(c.sent[0]))
		return errors.New(err)
	}
	if m.Message != msg {
		return errors.New(fmt.Sprintf("Expecting Error `%s`, found `%s`", msg, m.Message))
	}
	return nil
}

func TestSend(t *testing.T) {
	c := newTestClient()
	c.message(CommandMessage{"not-a-thing"})
	err := assertError(c, "Invalid command: 'not-a-thing'")
	if err != nil {
		t.Error(err)
	}
}

func TestBroadcast(t *testing.T) {
	hub.Broadcast(CommandMessage{"testing"})
	if len(testHub.sent) != 1 {
		t.Errorf("Expected 1 broadcasted message, found %d", len(testHub.sent))
	}
	c := testHub.sent[0].(CommandMessage)
	if c.Cmd != "testing" {
		t.Error("Expected command `testing`, found `%s`", c.Cmd)
	}
}
