package persistent

import (
	"log"
	"sync"
	"time"

	"github.com/doozr/guac/realtime"
	"github.com/doozr/jot"
)

// listen for incoming and outgoing events and pass them to the active connection.
//
// Terminate by closing the `done` channel.
func listen(r realtime.Connection, timeout time.Duration,
	receiveChan chan []byte, sendChan chan asyncEvent, done chan struct{}) {

	wg := sync.WaitGroup{}
	events := receive(r, done, &wg)
	send(r, sendChan, done, &wg)

	defer func() {
		jot.Print("persistent.listen: shutting down")

		// Close the connection to stop waiting websockets immediately
		r.Close()

		wg.Wait()
		jot.Print("persistent.listen: done")
	}()

	for {
		select {
		case <-done:
			return

		case event := <-events:
			receiveChan <- event

		case <-time.After(timeout):
			log.Printf("connection timeout after %s: reconnecting", timeout)
			return
		}
	}
}

// send outgoing messages.
func send(r realtime.Connection, sendChan chan asyncEvent,
	done chan struct{}, wg *sync.WaitGroup) {
	jot.Print("persistent.send: started")

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			jot.Print("persistent.send: done")
		}()

		for {
			select {
			case <-done:
				return

			case asyncEv := <-sendChan:
				err := r.Send(asyncEv.event)
				asyncEv.callback <- err
			}
		}
	}()
}

// receive incoming events from a connection and put them on a channel.
//
// Terminate by closing the `done` channel and waiting on `wg`.
func receive(r realtime.Connection, done chan struct{}, wg *sync.WaitGroup) (events chan []byte) {
	jot.Print("persistent.receive: started")
	events = make(chan []byte)

	wg.Add(1)
	go func() {
		defer func() {
			jot.Println("persistent.receive: shutting down")
			close(events)
			wg.Done()
			jot.Print("persistent.receive: done")
		}()

		for {
			select {
			case <-done:
				return
			default:
				event, err := r.Receive()
				if err != nil {
					jot.Print("persistent.receive: error while receiving events: ", err)
					return
				}
				events <- event
			}
		}
	}()
	return
}
