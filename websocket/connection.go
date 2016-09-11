package websocket

import "golang.org/x/net/websocket"

// Connection represents a connected websocket
type connection struct {
	id        string
	websocket *websocket.Conn
}

// ID returns the ID associated with the websocket
func (c connection) ID() string {
	return c.id
}

// Send a payload over the websocket
func (c connection) Send(payload []byte) (err error) {
	err = websocket.Message.Send(c.websocket, payload)
	return
}

// Receive a payload from the websocket
func (c connection) Receive() (payload []byte, err error) {
	err = websocket.Message.Receive(c.websocket, &payload)
	return
}
