package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/doozr/guac/web"
	"github.com/doozr/jot"

	"golang.org/x/net/websocket"
)

// New websocket Dialer.
func New(client web.Client) Dialer {
	return dialer{client}
}

// Dialer creates websocket connections to Slack.
type dialer struct {
	client web.Client
}

// Dial a websocket.
func (d dialer) Dial() (conn Connection, err error) {
	wsurl, id, name, err := d.getWebsocketURL()
	if err != nil {
		return
	}
	jot.Print("Retrieved identity: ", id, name)

	jot.Print("Dialing ", wsurl)
	ws, err := websocket.Dial(wsurl, "", "https://api.slack.com/")
	if err != nil {
		return
	}

	conn = connection{
		websocket: ws,
		id:        id,
		name:      name,
	}
	return
}

// getWebsocketURL gets the socket URL via the Web API.
func (d dialer) getWebsocketURL() (wsurl string, id string, name string, err error) {
	body, err := d.client.Get("rtm.start", nil)
	if err != nil {
		return
	}

	if !body.Success() {
		err = fmt.Errorf("Slack error: %s", body.Error())
		return
	}

	var respObj struct {
		URL  string `json:"url"`
		Self struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"self"`
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
	id = respObj.Self.ID
	name = respObj.Self.Name
	return
}
