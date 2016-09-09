package realtime

import (
	"encoding/json"

	"github.com/doozr/guac/slack"
)

// Connection represents an open Slack RealTime connection
type Connection struct {
	dialer slack.Dialer
	raw    slack.RawConnection
}

// Connect to the Slack RealTime API
func Connect(dialer slack.Dialer) (conn *Connection, err error) {
	rawConnection, err := dialer.Dial()
	if err != nil {
		return
	}

	conn = &Connection{
		dialer: dialer,
		raw:    rawConnection,
	}
	return
}

// Receive a Slack RealTimeEvent
func (c Connection) Receive() (event slack.RealTimeEvent, err error) {
	payload, err := c.raw.Receive()
	if err != nil {
		return
	}

	eventObj := RealTimeEvent{}
	err = json.Unmarshal(payload, &eventObj)
	eventObj.Raw = payload
	event = eventObj

	if err != nil {
		return
	}

	return
}

// Send a Slack RealTime event
func (c Connection) Send(event slack.RealTimeEvent) (err error) {
	return
}

// RealTimeEvent is a concrete implementation of slack.RealTimeEvent
type RealTimeEvent struct {
	EventType string `json:"type"`
	Raw       []byte
}

// Type of the event
func (r RealTimeEvent) Type() string {
	return r.EventType
}

// Payload of the event
func (r RealTimeEvent) Payload() []byte {
	return r.Raw
}
