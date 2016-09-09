package realtime

import (
	"encoding/json"

	"github.com/doozr/guac/slack"
)

// Connection represents an open Slack RealTime connection
type Connection struct {
	raw slack.RawConnection
}

// New to the Slack RealTime API
func New(raw slack.RawConnection) (conn *Connection) {
	conn = &Connection{
		raw: raw,
	}
	return
}

// Receive a Slack RealTimeEvent
func (c Connection) Receive() (event slack.RealTimeEvent, err error) {
	payload, err := c.raw.Receive()
	if err != nil {
		return
	}

	eventObj := realTimeEvent{}
	err = json.Unmarshal(payload, &eventObj)
	if err != nil {
		return
	}

	eventObj.Raw = payload
	event = eventObj
	return
}

// Send a Slack RealTime event
func (c Connection) Send(event slack.RealTimeEvent) (err error) {
	err = c.raw.Send(event.Payload())
	return
}
