package guac

import (
	"encoding/json"

	"github.com/doozr/guac/realtime"
)

func convertEvent(raw realtime.RawEvent) (event interface{}, err error) {
	switch raw.EventType() {
	case "pong":
		e := RealTimePingPong{}
		err = json.Unmarshal(raw.Payload(), &e)
		event = e

	case "message":
		e := RealTimeMessage{}
		err = json.Unmarshal(raw.Payload(), &e)
		event = e

	case "user_change":
		e := RealTimeUserChange{}
		err = json.Unmarshal(raw.Payload(), &e)
		event = e
	}

	if err != nil {
		event = nil
	}

	return
}
