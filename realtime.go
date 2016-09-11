package guac

import (
	"sync/atomic"
	"time"

	"github.com/doozr/guac/slack"
)

var counter uint64

// Get the next ID for this run
func nextID() uint64 {
	return atomic.AddUint64(&counter, 1)
}

// realTime is a concrete implementation of guac.RealTime
type realTime struct {
	connection slack.RealTimeConnection
}

// Receive an event from the slack RTM API
func (g realTime) Receive() (event slack.RealTimeEvent, err error) {
	return g.connection.Receive()
}

// PostMessage sends a chat message to the given channel
func (g realTime) PostMessage(channel, text string) (err error) {
	id := nextID()
	m := RealTimeMessage{
		EventType: "message",
		ID:        id,
		Channel:   channel,
		User:      "",
		Text:      text,
	}

	return g.connection.Send(eventWrapper{m.EventType, m})
}

// Ping sends a ping request
func (g realTime) Ping() (err error) {
	id := nextID()
	m := RealTimePing{
		EventType: "ping",
		ID:        id,
		Timestamp: time.Now(),
	}
	return g.connection.Send(eventWrapper{m.EventType, m})
}
