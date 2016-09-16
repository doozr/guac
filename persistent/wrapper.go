// Package persistent contains a wrapper for guac.realtime to enable automatic
// connection retries and reconnection on failure.
package persistent

import (
	"fmt"
	"log"
	"sync"

	"github.com/doozr/guac/realtime"
	"github.com/doozr/guac/web"
	"github.com/doozr/jot"
)

type reconnect struct {
	id          string
	name        string
	client      web.Client
	receiveChan chan []byte
	sendChan    chan asyncEvent
	done        chan struct{}
	wg          *sync.WaitGroup
}

type asyncEvent struct {
	event    []byte
	callback chan error
}

// New connection to the Slack RealTime API.
func New(client web.Client) (conn realtime.Connection) {
	reconn := &reconnect{
		client:      client,
		receiveChan: make(chan []byte),
		sendChan:    make(chan asyncEvent),
		done:        make(chan struct{}),
		wg:          &sync.WaitGroup{},
	}

	reconn.wg.Add(1)
	ready := make(chan struct{})
	go reconn.run(ready)
	<-ready

	conn = reconn
	return
}

// run the persistent listener.
//
// An infinite loop terminated only by closing the `done` channel. Open a
// real time connection and consume from it until it dies, and immediately
// create a new one. Repeat ad infinitum.
//
// Terminate by closing the `c.done` channel and waiting on `c.wg`.
func (c reconnect) run(ready chan struct{}) {
	jot.Print("persistent.run: started")
	defer func() {
		jot.Print("persistent.run: done")
		c.wg.Done()
	}()

	for {
		select {
		case <-c.done:
			return
		default:
		}

		jot.Print("persistent.run: connecting to Slack")
		r, ok := mustConnect(c.client, c.done)
		if ok {
			// If we are connected, set the ID and name
			c.id = r.ID()
			c.name = r.Name()
			log.Print("persistent.run: connected as ", r.Name())

			// Only close ready if it's open
			select {
			case <-ready:
			default:
				jot.Print("persistent.run: initial connection ready")
				close(ready)
			}

			jot.Print("persistent.run: listening for events")
			listen(r, c.receiveChan, c.sendChan, c.done)
		}
	}
}

// ID of the connected bot.
func (c reconnect) ID() string {
	return c.id
}

// Name of the connected bot.
func (c reconnect) Name() string {
	return c.name
}

// Close the persistent connection loop immediately.
func (c reconnect) Close() {
	jot.Print("persistent.wrapper: closing down connections")
	close(c.done)
	c.wg.Wait()

	jot.Print("persistent.wrapper: closing down internal channels")
	close(c.sendChan)
	close(c.receiveChan)
}

// Send an asynchronous event and wait for confirmation.
func (c reconnect) Send(raw []byte) (err error) {
	callback := make(chan error)
	a := asyncEvent{
		event:    raw,
		callback: callback,
	}

	jot.Print("persistent.Send sending async event: ", a)
	c.sendChan <- a

	jot.Print("persistent.Send awaiting callback")
	err = <-callback

	jot.Print("persistent.Send callback received")
	close(callback)
	return err
}

// Receive an incoming event.
func (c reconnect) Receive() (raw []byte, err error) {
	jot.Print("persistent.Receive: awaiting event")
	raw, ok := <-c.receiveChan
	if !ok {
		jot.Print("persistent.Receive: error")
		err = fmt.Errorf("Channel closed")
		return
	}

	jot.Print("persistent.Receive: event ", string(raw))
	return
}
