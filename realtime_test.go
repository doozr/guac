package guac

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/doozr/guac/realtime"
)

type TestRealTimeConnection struct {
	closed  chan bool
	receive func() (realtime.RawEvent, error)
	send    func(realtime.RawEvent) error
}

func (c TestRealTimeConnection) Close() {
	close(c.closed)
}

func (c TestRealTimeConnection) Send(event realtime.RawEvent) error {
	return c.send(event)
}

func (c TestRealTimeConnection) Receive() (realtime.RawEvent, error) {
	return c.receive()
}

type TestRealTimeEvent struct {
	eventType string
	payload   []byte
}

func (e TestRealTimeEvent) EventType() string {
	return e.eventType
}

func (e TestRealTimeEvent) Payload() []byte {
	return e.payload
}

func TestClosed(t *testing.T) {
	realTimeConnection := TestRealTimeConnection{
		closed: make(chan bool),
	}

	realTime := RealTimeClient{realTimeConnection}

	realTime.Close()

	select {
	case <-realTimeConnection.closed:
	default:
		t.Fatal("Expected connection to be closed")
	}
}

func receiveEvent(t *testing.T, eventType string, payload string, expected interface{}) {
	bytes := []byte(payload)
	called := false
	realTimeConnection := TestRealTimeConnection{
		receive: func() (realtime.RawEvent, error) {
			if called {
				t.Fatal("RealTimeConnection.Receive called more than once")
			}
			called = true
			return TestRealTimeEvent{
				eventType: eventType,
				payload:   bytes,
			}, nil
		},
	}

	realTime := RealTimeClient{realTimeConnection}

	event, err := realTime.Receive()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if !reflect.DeepEqual(expected, event) {
		t.Fatal("Event did not match", expected, event)
	}
}

func TestReceivePong(t *testing.T) {
	receiveEvent(t, "pong",
		`{
			"type": "pong",
			"id": 1234
		}`,
		PingPongEvent{
			Type: "pong",
			ID:   1234,
		})
}

func TestReceiveMessage(t *testing.T) {
	receiveEvent(t, "message",
		`{
			"type": "message",
			"id": 1234,
			"channel": "C9876543",
			"user": "U1234567",
			"text": "this is the text"
		}`,
		MessageEvent{
			Type:    "message",
			ID:      1234,
			Channel: "C9876543",
			User:    "U1234567",
			Text:    "this is the text",
		})
}

func TestReceiveUserChance(t *testing.T) {
	receiveEvent(t, "user_change",
		`{
			"type": "user_change",
			"user": {
				"id": "U1234567",
				"name": "Mr Test"
			}
		}`,
		UserChangeEvent{
			Type: "user_change",
			UserInfo: UserInfo{
				ID:   "U1234567",
				Name: "Mr Test",
			},
		})
}

func TestDoesNotReturnUnknown(t *testing.T) {
	type eventFn func() (string, []byte)
	incoming := make(chan eventFn, 2)
	incoming <- func() (string, []byte) { return "unknown", []byte(`{ "type": "uknown", "field": "value" }`) }
	incoming <- func() (string, []byte) { return "pong", []byte(`{ "type": "pong", "id": 1234 }`) }
	realTimeConnection := TestRealTimeConnection{
		receive: func() (realtime.RawEvent, error) {
			fn := <-incoming
			eventType, bytes := fn()
			return TestRealTimeEvent{
				eventType: eventType,
				payload:   bytes,
			}, nil
		},
	}

	realTime := RealTimeClient{realTimeConnection}

	event, err := realTime.Receive()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if _, ok := event.(PingPongEvent); !ok {
		t.Fatal("Expected RealTimePing instance", event)
	}
}

func TestReceiveError(t *testing.T) {
	realTimeConnection := TestRealTimeConnection{
		receive: func() (realtime.RawEvent, error) {
			return nil, fmt.Errorf("Receive error")
		},
	}

	rtm := RealTimeClient{realTimeConnection}

	event, err := rtm.Receive()

	if err == nil {
		t.Fatal("Expected error", err)
	}

	if event != nil {
		t.Fatal("Expected nil event", event)
	}
}

func TestPing(t *testing.T) {
	var event realtime.RawEvent
	realTimeConnection := TestRealTimeConnection{
		send: func(e realtime.RawEvent) error {
			if event != nil {
				t.Fatal("realTimeConnection.Send called more than once")
			}
			event = e
			return nil
		},
	}

	rtm := RealTimeClient{realTimeConnection}

	err := rtm.Ping()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if event == nil {
		t.Fatal("No event sent")
	}

	if event.EventType() != "ping" {
		t.Fatal("Event type should be `ping`", event.EventType())
	}

	var ping PingPongEvent
	err = json.Unmarshal(event.Payload(), &ping)

	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if ping.ID <= 0 {
		t.Fatal("ID should be > 0")
	}
}

func TestPingError(t *testing.T) {
	realTimeConnection := TestRealTimeConnection{
		send: func(e realtime.RawEvent) error {
			return fmt.Errorf("Ping error")
		},
	}

	rtm := RealTimeClient{realTimeConnection}

	err := rtm.Ping()
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestPostMessage(t *testing.T) {
	var event realtime.RawEvent
	realTimeConnection := TestRealTimeConnection{
		send: func(e realtime.RawEvent) error {
			if event != nil {
				t.Fatal("realTimeConnection.Send called more than once")
			}
			event = e
			return nil
		},
	}

	rtm := RealTimeClient{realTimeConnection}

	channel := "#F00DD00D"
	text := "this is the message"
	err := rtm.PostMessage(channel, text)
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if event == nil {
		t.Fatal("No event sent")
	}

	if event.EventType() != "message" {
		t.Fatal("Event type should be `message`", event.EventType())
	}

	var message MessageEvent
	err = json.Unmarshal(event.Payload(), &message)

	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if message.ID <= 0 {
		t.Fatal("ID should be > 0")
	}

	if message.Channel != channel {
		t.Fatal("Channel does not match", channel, message.Channel)
	}

	if message.Text != text {
		t.Fatal("Text does not match", text, message.Text)
	}

	if message.User != "" {
		t.Fatal("User should be blank", message)
	}
}

func TestPostMessageError(t *testing.T) {
	realTimeConnection := TestRealTimeConnection{
		send: func(e realtime.RawEvent) error {
			return fmt.Errorf("PostMessage error")
		},
	}

	rtm := RealTimeClient{realTimeConnection}

	channel := "#F00DD00D"
	text := "this is the message"
	err := rtm.PostMessage(channel, text)
	if err == nil {
		t.Fatal("Expected error")
	}
}
