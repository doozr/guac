package websocket

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
