package guac

// RealTimeMessage is a chat message sent to a user or channel
type RealTimeMessage struct {
	EventType string `json:"type"`
	ID        uint64 `json:"id"`
	Channel   string `json:"channel"`
	User      string `json:"user"`
	Text      string `json:"text"`
}

// RealTimePing is a ping and also the reciprocal pong
type RealTimePing struct {
	EventType string `json:"type"`
	ID        uint64 `json:"id"`
}
