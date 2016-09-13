package guac

import (
	"encoding/json"

	"github.com/doozr/guac/realtime"
)

// convertEvent accepts any RawEvent and converts it to a concrete type.
func convertEvent(raw realtime.RawEvent) (event interface{}, err error) {
	switch raw.EventType() {
	case "pong":
		e := PingPongEvent{}
		err = json.Unmarshal(raw.Payload(), &e)
		event = e

	case "message":
		e := MessageEvent{}
		err = json.Unmarshal(raw.Payload(), &e)
		event = e

	case "user_change":
		e := UserChangeEvent{}
		err = json.Unmarshal(raw.Payload(), &e)
		event = e
	}

	if err != nil {
		event = nil
	}

	return
}
