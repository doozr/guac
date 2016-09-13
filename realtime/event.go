package realtime

import "encoding/json"

// realTimeEvent is a concrete implementation of Event.
type realTimeEvent struct {
	eventType string
	payload   []byte
}

// UnmarshalJSON reads the type from the JSON and stores the raw payload.
func (r *realTimeEvent) UnmarshalJSON(payload []byte) (err error) {
	r.payload = payload

	var parsed interface{}
	err = json.Unmarshal(payload, &parsed)
	if err != nil {
		return
	}

	if m, ok := parsed.(map[string]interface{}); ok {
		if v, ok := m["type"]; ok {
			if t, ok := v.(string); ok {
				r.eventType = t
			}
		}
	}

	return
}

// MarshalJSON returns the raw JSON representation of the event.
func (r *realTimeEvent) MarshalJSON() (payload []byte, err error) {
	payload = r.payload
	return
}

// Type of the event.
func (r realTimeEvent) EventType() string {
	return r.eventType
}

// Payload of the event.
func (r realTimeEvent) Payload() []byte {
	return r.payload
}
