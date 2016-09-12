package realtime

// Event is an incoming RealTime payload
type Event interface {
	EventType() string
	Payload() []byte
}

// Connection is an active Slack RealTime API connection
type Connection interface {
	Send(Event) error
	Receive() (Event, error)
}
