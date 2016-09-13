package realtime

import (
	"encoding/json"

	"github.com/doozr/guac/websocket"
)

// Connection represents an open Slack RealTime connection
type connection struct {
	raw websocket.Connection
}

// New to the Slack RealTime API
func New(raw websocket.Connection) (conn Connection) {
	conn = &connection{
		raw: raw,
	}
	return
}

func (c connection) Close() {
	c.raw.Close()
}

// Receive a Slack RealTimeEvent
func (c connection) Receive() (event RawEvent, err error) {
	payload, err := c.raw.Receive()
	if err != nil {
		return
	}

	eventObj := realTimeEvent{}
	err = json.Unmarshal(payload, &eventObj)
	if err != nil {
		return
	}

	eventObj.payload = payload
	event = eventObj
	return
}

// Send a Slack RealTime event
func (c connection) Send(event RawEvent) (err error) {
	err = c.raw.Send(event.Payload())
	return
}
