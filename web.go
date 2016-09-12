package guac

import (
	"github.com/doozr/guac/realtime"
	"github.com/doozr/guac/web"
	"github.com/doozr/guac/websocket"
)

// WebClient is an interface to the Slack Web API
type WebClient struct {
	client web.Client
}

// RealTime connects to the Slack RealTime API using the Web client's credentials
func (g WebClient) RealTime() (client RealTimeClient, err error) {
	raw, err := websocket.New(g.client).Dial()
	if err != nil {
		return
	}

	realTimeConn := realtime.New(raw)
	client = RealTimeClient{
		connection: realTimeConn,
	}
	return
}
