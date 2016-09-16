package persistent

import (
	"log"
	"time"

	"github.com/doozr/guac/realtime"
	"github.com/doozr/guac/web"
	"github.com/doozr/jot"
)

// mustConnect gets a real time connection to Slack and retries infinitely on failure.
//
// Backs off retries from 1s, 2s, 5s, 10s, 30s and 1 minute, and retries at 1 minute
// intervals after that. Close the `done` channel to stop retrying.
func mustConnect(client web.Client, done chan struct{}) (ws realtime.Connection, ok bool) {
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
	for {
		ws, err = realtime.New(client).Dial()
		if err == nil {
			jot.Print("persistent.mustConnect: websocket connected")
			ok = true
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
