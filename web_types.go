package guac

// UserInfo represents a single user's profile info.
type UserInfo struct {
	ID   string
	Name string
}

// ChannelInfo represents a channel or group's info.
type ChannelInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsMember bool   `json:"is_member"`
}
