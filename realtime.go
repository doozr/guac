package guac

import (
	"encoding/json"
	"sync/atomic"

	"github.com/doozr/guac/realtime"
	"github.com/doozr/jot"
)

var counter uint64

// nextID returns the next message ID for this run.
func nextID() uint64 {
	return atomic.AddUint64(&counter, 1)
}

// realTimeClient is the default implementation of RealTimeClient
type realTimeClient struct {
	WebClient
	connection realtime.Connection
}

// ID of the bot
func (c *realTimeClient) ID() string {
	return c.connection.ID()
}

// Name of the bot
func (c *realTimeClient) Name() string {
	return c.connection.Name()
}

func (c *realTimeClient) RealTime() (RealTimeClient, error) {
	return c, nil
}

// Close terminates the connection.
func (c *realTimeClient) Close() {
	c.connection.Close()
}

// PostMessage sends a chat message to the given channel.
//
// The message is posted as the bot itself, and does not try to take on the
// identity of a user. Use the API formatting standard.
func (c *realTimeClient) PostMessage(channel, text string) (err error) {
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
		err = c.connection.Send(payload)
	}
	return
}

// Ping sends a ping request.
//
// Sends a bare ping with no additional information.
func (c *realTimeClient) Ping() (err error) {
	id := nextID()
	m := PingPongEvent{
		Type: "ping",
		ID:   id,
	}
	payload, err := json.Marshal(m)
	if err == nil {
		err = c.connection.Send(payload)
	}
	return
}

// Receive an event or an error.
//
// Blocks until an known event type arrives or an error occurs.
func (c *realTimeClient) Receive() (event interface{}, err error) {
	var raw []byte
	for {
		// Bail out on error immediately
		raw, err = c.connection.Receive()
		if err != nil {
			jot.Print("realtimeclient.Receive error from realtime.Receive: ", err)
			return
		}
		jot.Print("realtime.Receive raw ", string(raw))

		// Only return if there is something worth returning, or an error
		event, err = convertEvent(raw)
		if err != nil {
			jot.Print("realtimeclient.Receive error converting event: ", err)
			return
		}

		if event != nil {
			jot.Print("realtime.Receive event ", event)
			return
		}
	}
}
