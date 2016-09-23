package guac

// EventChan is a channel of any type of event
type EventChan chan interface{}

// MessageEvent is a chat message sent to a user or channel.
type MessageEvent struct {
	Type    string `json:"type"`
	ID      uint64 `json:"id"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
}

// PingPongEvent is a ping and also the reciprocal pong.
type PingPongEvent struct {
	Type string `json:"type"`
	ID   uint64 `json:"id"`
}

// UserChangeEvent is a notification that something about a user has changed.
// Currently only username changes are supported.
type UserChangeEvent struct {
	Type     string `json:"type"`
	UserInfo `json:"user"`
}
