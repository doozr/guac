// Package web contains request formatting and response parsing for Slack
// Web API calls.
package web

import "net/url"

// Response from the Slack Web API.
type Response interface {
	Success() bool
	Error() error
	Payload() []byte
}

// Client is a client to the Slack Web API.
type Client interface {
	Get(string, url.Values) (Response, error)
	Post(string, url.Values) (Response, error)
}
