package reconnect

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
	receiveChan chan realtime.RawEvent
	sendChan    chan asyncEvent
	done        chan struct{}
	wg          *sync.WaitGroup
}

type asyncEvent struct {
	event    realtime.RawEvent
	callback chan error
}

// New connection to the Slack RealTime API.
func New(client web.Client) (conn realtime.Connection) {
	reconn := &reconnect{
		client:      client,
		receiveChan: make(chan realtime.RawEvent),
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

func (c reconnect) run(ready chan struct{}) {
	jot.Print("reconnect wrapper started")
	defer func() {
		jot.Print("reconnect wrapper done")
		c.wg.Done()
	}()

	for {
		select {
		case <-c.done:
			return
		default:
		}

		jot.Print("Connecting to Slack")
		r, ok := mustConnect(c.client, c.done)
		if ok {
			// If we are connected, set the ID and name
			c.id = r.ID()
			c.name = r.Name()
			log.Print("Connected as ", r.Name())

			// Only close ready if it's open
			select {
			case <-ready:
			default:
				jot.Print("initial reconnect ready")
				close(ready)
			}

			jot.Print("Listening for events")
			listen(r, c.receiveChan, c.sendChan, c.done, c.wg)

			jot.Print("Closing connection")
			r.Close()
		}
	}
}

func (c reconnect) ID() string {
	return c.id
}

func (c reconnect) Name() string {
	return c.name
}

func (c reconnect) Close() {
	close(c.done)
	close(c.sendChan)
	close(c.receiveChan)
	c.wg.Wait()
}

func (c reconnect) Send(event realtime.RawEvent) (err error) {
	callback := make(chan error)
	a := asyncEvent{
		event:    event,
		callback: callback,
	}

	jot.Print("reconnect.Send sending async event: ", a)
	c.sendChan <- a

	jot.Print("reconnect.Send awaiting callback")
	err = <-callback

	jot.Print("reconnect.Send callback received")
	close(callback)
	return err
}

func (c reconnect) Receive() (event realtime.RawEvent, err error) {
	jot.Print("reconnect.Receive awaiting event")
	event, ok := <-c.receiveChan
	if !ok {
		jot.Print("reconnect.Receive error")
		err = fmt.Errorf("Channel closed")
		return
	}

	jot.Print("reconnect.Receive event ", event.EventType(), " ", string(event.Payload()))
	return
}
