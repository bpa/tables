package test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/bpa/tables/data"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

type Client struct {
	Sent          []interface{}
	Player        *data.Player
	handleMessage func(data.Client, []byte)
}

func NewClient(f func(data.Client, []byte)) *Client {
	return &Client{
		Sent:          make([]interface{}, 0, 3),
		handleMessage: f,
	}
}

func Ignore(_ data.Client, _ []byte) {}

func (tc *Client) Error(err string, params ...interface{}) {
	tc.Sent = append(tc.Sent, ErrorMessage{
		fmt.Sprintf(err, params...),
	})
}

func (tc *Client) Send(v interface{}) {
	tc.Sent = append(tc.Sent, v)
}

func (tc *Client) Broadcast(v interface{}) {
	tc.Sent = append(tc.Sent, v)
}

func (tc *Client) Message(v interface{}) {
	g, _ := json.Marshal(v)
	tc.handleMessage(tc, g)
}

func (tc *Client) GetPlayer() *data.Player {
	return tc.Player
}

func (tc *Client) SetPlayer(p *data.Player) {
	tc.Player = p
}

func (tc *Client) Host() string {
	return "localhost"
}

func (tc *Client) setPlayer(player *data.Player) {
	tc.Player = player
}

func (tc *Client) AssertError(t *testing.T, msg string) {
	t.Helper()
	if len(tc.Sent) < 1 {
		t.Errorf("Expecting Error: `%s`, found nothing", msg)
		return
	}
	m, ok := tc.Sent[0].(ErrorMessage)
	if ok {
		if m.Error != msg {
			t.Errorf("Expecting Error `%s`, found `%s`", msg, m.Error)
		}
	} else {
		t.Errorf("Expecting CommandMessage, found `%s`", reflect.TypeOf(tc.Sent[0]))
	}
}
