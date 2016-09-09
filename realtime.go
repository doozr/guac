package guac

import "github.com/doozr/guac/slack"

type realTime struct {
	connection slack.RealTimeConnection
}

func (g realTime) GetEvent() (event slack.RealTimeEvent, err error) {
	return g.connection.Receive()
}

func (g realTime) PostMessage(channel, text string) (err error) {
	return
}
