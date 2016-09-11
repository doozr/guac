package websocket

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"

	"github.com/doozr/guac/slack"
)

// New websocket dialer
func New(client slack.WebClient) slack.Dialer {
	return dialer{client}
}

// Dialer creates websocket connections to Slack
type dialer struct {
	client slack.WebClient
}

// Dial a websocket
func (d dialer) Dial() (conn slack.RawConnection, err error) {
	wsurl, id, err := d.getWebsocketURL()
	if err != nil {
		return
	}

	ws, err := websocket.Dial(wsurl, "", "https://api.slack.com/")
	if err != nil {
		return
	}

	conn = connection{
		websocket: ws,
		id:        id,
	}
	return
}

// getWebsocketURL via the Web API
func (d dialer) getWebsocketURL() (wsurl string, id string, err error) {
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
