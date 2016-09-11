package guac

import (
	"github.com/doozr/guac/realtime"
	"github.com/doozr/guac/slack"
	"github.com/doozr/guac/websocket"
)

type guac struct {
	client slack.WebClient
}

func (g guac) RealTime() (conn RealTime, err error) {
	raw, err := websocket.New(g.client).Dial()
	if err != nil {
		return
	}

	realTimeConn := realtime.New(raw)
	conn = realTime{
		connection: realTimeConn,
	}
	return
}
