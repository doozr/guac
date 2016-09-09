package guac

import "encoding/json"

// Wrapper to make any event satisfy the slack.RealTimeEvent interface
type eventWrapper struct {
	EventType string
	Event     interface{}
}

func (w eventWrapper) Type() string {
	return w.EventType
}

func (w eventWrapper) Payload() []byte {
	payload, _ := json.Marshal(w.Event)
	return payload
}
