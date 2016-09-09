package realtime

// realTimeEvent is a concrete implementation of slack.RealTimeEvent
type realTimeEvent struct {
	EventType string `json:"type"`
	Raw       []byte
}

// Type of the event
func (r realTimeEvent) Type() string {
	return r.EventType
}

// Payload of the event
func (r realTimeEvent) Payload() []byte {
	return r.Raw
}
