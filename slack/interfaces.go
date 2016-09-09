package slack

import "net/url"

// RawConnection sends and receives byte arrays to the Slack RealTime API
type RawConnection interface {
	ID() string
	Send([]byte) error
	Receive() ([]byte, error)
}

// Dialer makes raw connections to the Slack RealTime API
type Dialer interface {
	Dial() (RawConnection, error)
}

// RealTimeEvent is an incoming RealTime payload
type RealTimeEvent interface {
	Type() string
	Payload() []byte
}

// RealTimeConnection is an active Slack RealTime API connection
type RealTimeConnection interface {
	Send(RealTimeEvent) error
	Receive() (RealTimeEvent, error)
}

// APIResponse from the Slack Web API
type APIResponse interface {
	Success() bool
	Error() error
	Payload() []byte
}

// WebClient is a client to the Slack Web API
type WebClient interface {
	Get(endPoint string, values url.Values) (APIResponse, error)
	Post(endPoint string, values url.Values) (APIResponse, error)
}
