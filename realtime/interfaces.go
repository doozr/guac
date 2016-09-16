// Package realtime contains a websocket dialer and connection for sending
// and receiving JSON as raw bytes from the Slack API.
package realtime

// Connection sends and receives byte arrays to the Slack RealTime API.
type Connection interface {
	Close()
	ID() string
	Name() string
	Send([]byte) error
	Receive() ([]byte, error)
}

// Dialer makes raw connections to the Slack RealTime API.
type Dialer interface {
	Dial() (Connection, error)
}
