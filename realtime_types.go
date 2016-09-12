package guac

// RealTimeMessage is a chat message sent to a user or channel
type RealTimeMessage struct {
	Type    string `json:"type"`
	ID      uint64 `json:"id"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
}

// RealTimePing is a ping and also the reciprocal pong
type RealTimePing struct {
	Type string `json:"type"`
	ID   uint64 `json:"id"`
}

// RealTimeUserChange is a notification that something about a user has changed
// Currently only username changes are supported
type RealTimeUserChange struct {
	Type string   `json:"type"`
	User UserInfo `json:"user"`
}
