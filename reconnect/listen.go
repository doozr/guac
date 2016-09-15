package reconnect

import (
	"sync"

	"github.com/doozr/guac/realtime"
	"github.com/doozr/jot"
)

func listen(r realtime.Connection,
	receiveChan chan realtime.RawEvent, sendChan chan asyncEvent,
	done chan struct{}, wg *sync.WaitGroup) {

	events := receive(r, done, wg)

	for {
		select {
		case <-done:
			return
		case event := <-events:
			receiveChan <- event
		case asyncEv, ok := <-sendChan:
			if !ok {
				return
			}
			err := r.Send(asyncEv.event)
			asyncEv.callback <- err
		}
	}
}

func receive(r realtime.Connection, done chan struct{}, wg *sync.WaitGroup) (events chan realtime.RawEvent) {
	jot.Print("reconnect receive started")
	events = make(chan realtime.RawEvent)
	wg.Add(1)

	go func() {
		defer func() {
			close(events)
			wg.Done()
			jot.Print("reconnect receive done")
		}()

		for {
			select {
			case <-done:
				jot.Println("Terminating listener")
				return
			default:
				event, err := r.Receive()
				if err != nil {
					jot.Print("Error while receiving events: ", err)
					return
				}
				events <- event
			}
		}
	}()
	return
}
