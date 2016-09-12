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

// PostMessage sends a chat message to the given channel
func (g RealTimeClient) PostMessage(channel, text string) (err error) {
	id := nextID()
	m := RealTimeMessage{
		ID:      id,
		Channel: channel,
		User:    "",
		Text:    text,
	}

	return g.connection.Send(eventWrapper{"message", m})
}

// Ping sends a ping request
func (g RealTimeClient) Ping() (err error) {
	id := nextID()
	m := RealTimePing{
		ID: id,
	}
	return g.connection.Send(eventWrapper{"ping", m})
}

// Receive an event from the Slack RealTime API
func (g RealTimeClient) Receive() (event interface{}, err error) {
	var raw realtime.RawEvent
	for {
		// Bail out on error immediately
		raw, err = g.connection.Receive()
		if err != nil {
			break
		}

		// Only return if there is something worth returning, or an error
		event, err = convertEvent(raw)
		if event != nil || err != nil {
			break
		}
	}
	return
}
