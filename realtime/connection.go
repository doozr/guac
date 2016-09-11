package realtime

import (
	"encoding/json"

	"github.com/doozr/guac/slack"
)

// Connection represents an open Slack RealTime connection
type connection struct {
	raw slack.RawConnection
}

// New to the Slack RealTime API
func New(raw slack.RawConnection) (conn slack.RealTimeConnection) {
	conn = &connection{
		raw: raw,
	}
	return
}

// Receive a Slack RealTimeEvent
func (c connection) Receive() (event slack.RealTimeEvent, err error) {
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
func (c connection) Send(event slack.RealTimeEvent) (err error) {
	err = c.raw.Send(event.Payload())
	return
}
