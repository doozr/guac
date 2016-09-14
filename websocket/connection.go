package websocket

import "golang.org/x/net/websocket"

// Connection represents a connected websocket.
type connection struct {
	id        string
	name      string
	websocket *websocket.Conn
	err       error
}

// Close terminates the web socket.
func (c connection) Close() {
	c.websocket.Close()
}

// ID returns the ID associated with the websocket.
func (c connection) ID() string {
	return c.id
}

// Name returns the username associated with the websocket
func (c connection) Name() string {
	return c.name
}

// Send a payload over the websocket.
func (c connection) Send(payload []byte) (err error) {
	if c.err != nil {
		return c.err
	}

	// Force it to be a TextFrame by wrapping in string
	err = websocket.Message.Send(c.websocket, string(payload))
	if err != nil {
		c.err = err
	}
	return
}

// Receive a payload from the websocket.
func (c connection) Receive() (payload []byte, err error) {
	if c.err != nil {
		return nil, c.err
	}

	err = websocket.Message.Receive(c.websocket, &payload)
	if err != nil {
		c.err = err
	}
	return
}
