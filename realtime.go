package guac

import (
	"sync/atomic"

	"github.com/doozr/guac/realtime"
)

var counter uint64 = 1

// Get the next ID for this run
func nextID() uint64 {
	return atomic.AddUint64(&counter, 1)
}

// RealTimeClient is a client of the Slack RealTime API
type RealTimeClient struct {
	connection realtime.Connection
}

// Receive an event from the Slack RealTime API
func (g RealTimeClient) Receive() (event realtime.Event, err error) {
	return g.connection.Receive()
}

// PostMessage sends a chat message to the given channel
func (g RealTimeClient) PostMessage(channel, text string) (err error) {
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
func (g RealTimeClient) Ping() (err error) {
	id := nextID()
	m := RealTimePing{
		EventType: "ping",
		ID:        id,
	}
	return g.connection.Send(eventWrapper{m.EventType, m})
}
