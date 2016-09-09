package websocket

import "golang.org/x/net/websocket"

// Connection represents a connected websocket
type Connection struct {
	id        string
	websocket *websocket.Conn
}

// ID returns the ID associated with the websocket
func (c Connection) ID() string {
	return c.id
}

// Send a payload over the websocket
func (c Connection) Send(payload []byte) (err error) {
	err = websocket.Message.Send(c.websocket, payload)
	return
}

// Receive a payload from the websocket
func (c Connection) Receive() (payload []byte, err error) {
	err = websocket.Message.Receive(c.websocket, &payload)
	return
}
