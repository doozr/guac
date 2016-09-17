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

// RealTimeClient is a client of the Slack RealTime API.
//
// The connection stays open between calls until Close is called. If an error
// is returned at any point, it should be considered fatal for the connection
// and a new connection should be opened with WebClient.RealTime.
//
// Subsequent calls after an error will result in the same error.
type RealTimeClient struct {
	WebClient
	connection  realtime.Connection
	receiveChan chan interface{}
}

// ID of the bot
func (c *RealTimeClient) ID() string {
	return c.connection.ID()
}

// Name of the bot
func (c *RealTimeClient) Name() string {
	return c.connection.Name()
}

// Close terminates the connection.
func (c *RealTimeClient) Close() {
	c.connection.Close()
}

// PostMessage sends a chat message to the given channel.
//
// The message is posted as the bot itself, and does not try to take on the
// identity of a user. Use the API formatting standard.
func (c *RealTimeClient) PostMessage(channel, text string) (err error) {
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
func (c *RealTimeClient) Ping() (err error) {
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

//Receive an event from the Slack RealTime API.
//
//Receive one of the concrete event types and return it. The event should be
//checked with a type assertion to determine its type. If a message of an
//as-yet unsupported type arrives it will be ignored.
func (c *RealTimeClient) Receive() chan interface{} {
	if c.receiveChan != nil {
		jot.Print("realtimeclient.Receive already running")
		return c.receiveChan
	}

	jot.Print("realtimeclient.Receive started")
	c.receiveChan = make(chan interface{})

	go func() {
		defer func() {
			close(c.receiveChan)
			c.receiveChan = nil
			jot.Print("realtime.Receive done")
		}()

		for {
			// Bail out on error immediately
			raw, err := c.connection.Receive()
			if err != nil {
				jot.Print("realtimeclient.Receive error from realtime.Receive: ", err)
				return
			}

			// Only return if there is something worth returning, or an error
			event, err := convertEvent(raw)
			if err != nil {
				jot.Print("realtimeclient.Receive error converting event: ", err)
				return
			}

			//
			if event != nil {
				c.receiveChan <- event
			}
		}
	}()
	return c.receiveChan
}
