package realtime

// Event is the base interface to any Slack RealTime event.
type Event interface {
	EventType() string
}

// RawEvent is a raw Slack RealTime event ready for receiving or sending.
type RawEvent interface {
	EventType() string
	Payload() []byte
}

// Connection is an active Slack RealTime API connection.
type Connection interface {
	ID() string
	Close()
	Send(RawEvent) error
	Receive() (RawEvent, error)
}
