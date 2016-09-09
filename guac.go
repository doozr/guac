package guac

import (
	"github.com/doozr/guac/slack"
	"github.com/doozr/guac/web"
)

// Guac interface to Slack
type Guac interface {
	RealTime() (RealTime, error)
}

// RealTime interface to Slack RealTime
type RealTime interface {
	Receive() (slack.RealTimeEvent, error)
	PostMessage(channel, text string) error
	Ping() error
}

// New Guac instance
func New(token string) Guac {
	return guac{
		client: web.New(token),
	}
}
