package guac

import "encoding/json"

// Wrapper to make any event satisfy the slack.RealTimeEvent interface
type eventWrapper struct {
	eventType string
	event     interface{}
}

func (w eventWrapper) EventType() string {
	return w.eventType
}

func (w eventWrapper) Payload() []byte {
	payload, _ := json.Marshal(w.event)
	return payload
}
