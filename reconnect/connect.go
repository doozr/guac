package reconnect

import (
	"log"
	"time"

	"github.com/doozr/guac/realtime"
	"github.com/doozr/guac/web"
	"github.com/doozr/guac/websocket"
)

func mustConnect(client web.Client, done chan struct{}) (r realtime.Connection, ok bool) {
	backoffTimes := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		5 * time.Second,
		10 * time.Second,
		30 * time.Second,
		60 * time.Second,
	}
	var backoff time.Duration

	var err error
	var ws websocket.Connection
	for {
		ws, err = websocket.New(client).Dial()
		if err == nil {
			ok = true
			r = realtime.New(ws)
			return
		}

		// If we can increase backoff time, do so, otherwise stick with
		// whatever value we have (last value of backoffTimes)
		if len(backoffTimes) > 0 {
			backoff, backoffTimes = backoffTimes[0], backoffTimes[1:]
		}

		log.Printf("Error while connecting; retrying in %s: %v", backoff, err)

		// Wait for either done or sleep to complete
		select {
		case <-done:
			return
		case <-time.After(backoff):
		}
	}
}
