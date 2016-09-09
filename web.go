package guac

import (
	"github.com/doozr/guac/realtime"
	"github.com/doozr/guac/web"
	"github.com/doozr/guac/websocket"
)

type guac struct {
	client web.Client
}

func (g guac) RealTime() (conn RealTime, err error) {
	dialer := websocket.New(g.client)
	RealTimeConn, err := realtime.Connect(dialer)
	if err != nil {
		return
	}

	conn = realTime{
		connection: RealTimeConn,
	}
	return
}
