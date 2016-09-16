package guac

import (
	"encoding/json"
	"sync/atomic"

	"github.com/doozr/guac/realtime"
)

var counter uint64

// nextID returns the next message ID for this run.
func nextID() uint64 {
	return atomic.AddUint64(&counter, 1)
}

// RealTimeClient is a client of the Slack RealTime API.
//
// The connection stays open between calls until Close is called. If an error
// is returned at any point, it should be considered fatal for the connection
// and a new connection should be opened with WebClient.RealTime.
//
// Subsequent calls after an error will result in the same error.
type RealTimeClient struct {
	WebClient
	connection realtime.Connection
}

// ID of the bot
func (g RealTimeClient) ID() string {
	return g.connection.ID()
}

// Name of the bot
func (g RealTimeClient) Name() string {
	return g.connection.Name()
}

// Close terminates the connection.
func (g RealTimeClient) Close() {
	g.connection.Close()
}

// PostMessage sends a chat message to the given channel.
//
// The message is posted as the bot itself, and does not try to take on the
// identity of a user. Use the API formatting standard.
func (g RealTimeClient) PostMessage(channel, text string) (err error) {
	id := nextID()
	m := MessageEvent{
		Type:    "message",
		ID:      id,
		Channel: channel,
		User:    "",
		Text:    text,
	}
	payload, err := json.Marshal(m)
	if err == nil {
		err = g.connection.Send(payload)
	}
	return
}

// Ping sends a ping request.
//
// Sends a bare ping with no additional information.
func (g RealTimeClient) Ping() (err error) {
	id := nextID()
	m := PingPongEvent{
		Type: "ping",
		ID:   id,
	}
	payload, err := json.Marshal(m)
	if err == nil {
		err = g.connection.Send(payload)
	}
	return
}

//Receive an event from the Slack RealTime API.
//
//Receive one of the concrete event types and return it. The event should be
//checked with a type assertion to determine its type. If a message of an
//as-yet unsupported type arrives it will be ignored.
func (g RealTimeClient) Receive() (event interface{}, err error) {
	var raw []byte
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
