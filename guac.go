// Package guac contains a client for the Slack Web and Real Time APIs in Go.
package guac

import (
	"net/http"
	"strings"

	"github.com/doozr/guac/web"
)

// WebClient is an interface to the Slack Web API.
type WebClient interface {
	RealTime() (RealTimeClient, error)
	UsersList() ([]UserInfo, error)
	ChannelsList() ([]ChannelInfo, error)
	GroupsList() ([]ChannelInfo, error)
	IMOpen(string) (string, error)
}

// RealTimeClient is a client of the Slack RealTime API.
type RealTimeClient interface {
	WebClient
	ID() string
	Name() string
	Close()
	PostMessage(string, string) error
	Ping() error
	Receive() (interface{}, error)
}

// New Slack Web API client.
func New(token string) WebClient {
	return &webClient{
		client: web.New(token, &http.Client{}),
	}
}

type stringEncoding struct {
	raw    string
	cooked string
}

var stringEncodings = []stringEncoding{
	{"&", "&amp;"},
	{"<", "&lt;"},
	{">", "&gt;"},
}

// EncodingString performs limited URL encoding as per Slack encoding standards
func EncodeString(input string) string {
	output := input
	for _, encoding := range stringEncodings {
		output = strings.Replace(output, encoding.raw, encoding.cooked, -1)
	}
	return output
}

// DecodeString performs limited URL decoding as per Slack encoding standards
func DecodeString(input string) string {
	output := input
	for _, encoding := range stringEncodings {
		output = strings.Replace(output, encoding.cooked, encoding.raw, -1)
	}
	return output
}
