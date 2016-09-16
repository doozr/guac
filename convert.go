package guac

import "encoding/json"

type basicEvent struct {
	Type string `json:"type"`
}

// convertEvent accepts any raw bytes and converts them to a concrete type.
func convertEvent(raw []byte) (event interface{}, err error) {
	b := basicEvent{}
	err = json.Unmarshal(raw, &b)
	if err != nil {
		return
	}

	switch b.Type {
	case "pong":
		e := PingPongEvent{}
		err = json.Unmarshal(raw, &e)
		event = e

	case "message":
		e := MessageEvent{}
		err = json.Unmarshal(raw, &e)
		event = e

	case "user_change":
		e := UserChangeEvent{}
		err = json.Unmarshal(raw, &e)
		event = e
	}

	if err != nil {
		event = nil
	}

	return
}
