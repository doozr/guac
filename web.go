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
