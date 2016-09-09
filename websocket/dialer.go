package websocket

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"

	"github.com/doozr/guac/slack"
)

// New websocket dialer
func New(client slack.WebClient) slack.Dialer {
	return Dialer{client}
}

// Connection represents a connected websocket
type Connection struct {
	id        string
	websocket *websocket.Conn
}

// ID returns the ID associated with the websocket
func (c Connection) ID() string {
	return c.id
}

// Send a payload over the websocket
func (c Connection) Send(payload []byte) (err error) {
	err = websocket.Message.Send(c.websocket, payload)
	return
}

// Receive a payload from the websocket
func (c Connection) Receive() (payload []byte, err error) {
	err = websocket.Message.Receive(c.websocket, &payload)
	return
}

// Dialer creates websocket connections to Slack
type Dialer struct {
	client slack.WebClient
}

// Dial a websocket
func (d Dialer) Dial() (conn slack.RawConnection, err error) {
	wsurl, id, err := d.getWebsocketURL()
	if err != nil {
		return
	}

	ws, err := websocket.Dial(wsurl, "", "https://api.slack.com/")
	if err != nil {
		return
	}

	conn = Connection{
		websocket: ws,
		id:        id,
	}
	return
}

func (d Dialer) getWebsocketURL() (wsurl string, id string, err error) {
	body, err := d.client.Get("rtm.start", nil)
	if err != nil {
		return
	}

	if !body.Success() {
		err = fmt.Errorf("Slack error: %s", body.Error())
		return
	}

	var respObj struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	}

	err = json.Unmarshal(body.Payload(), &respObj)
	if err != nil {
		return
	}

	if respObj.URL == "" {
		err = fmt.Errorf("No websocket URL received")
		return
	}

	wsurl = respObj.URL
	id = respObj.ID
	return
}
