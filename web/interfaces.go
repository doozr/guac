package web

import "net/url"

// Response from the Slack Web API
type Response interface {
	Success() bool
	Error() error
	Payload() []byte
}

// Client is a client to the Slack Web API
type Client interface {
	Get(endPoint string, values url.Values) (Response, error)
	Post(endPoint string, values url.Values) (Response, error)
}
