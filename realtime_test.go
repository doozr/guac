package guac

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/doozr/guac/realtime"
)

type TestRealTimeConnection struct {
	receive func() (realtime.Event, error)
	send    func(realtime.Event) error
}

func (c TestRealTimeConnection) Send(event realtime.Event) error {
	return c.send(event)
}

func (c TestRealTimeConnection) Receive() (realtime.Event, error) {
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

func TestSuccessfulReceive(t *testing.T) {
	payload := []byte(`this is a test payload`)
	called := false
	realTimeConnection := TestRealTimeConnection{
		receive: func() (realtime.Event, error) {
			if called {
				t.Fatal("RealTimeConnection.Receive called more than once")
			}
			called = true
			return TestRealTimeEvent{
				eventType: "test",
				payload:   payload,
			}, nil
		},
	}

	realTime := realTime{realTimeConnection}

	event, err := realTime.Receive()

	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if event.EventType() != "test" {
		t.Fatal("Event type does not match", "test", event.EventType())
	}

	if !reflect.DeepEqual(event.Payload(), payload) {
		t.Fatal("Payload does not match", payload, event.Payload())
	}
}

func TestReceiveError(t *testing.T) {
	realTimeConnection := TestRealTimeConnection{
		receive: func() (realtime.Event, error) {
			return nil, fmt.Errorf("Receive error")
		},
	}

	rtm := realTime{realTimeConnection}

	event, err := rtm.Receive()

	if err == nil {
		t.Fatal("Expected error", err)
	}

	if event != nil {
		t.Fatal("Expected nil event", event)
	}
}

func TestPing(t *testing.T) {
	var event realtime.Event
	realTimeConnection := TestRealTimeConnection{
		send: func(e realtime.Event) error {
			if event != nil {
				t.Fatal("realTimeConnection.Send called more than once")
			}
			event = e
			return nil
		},
	}

	rtm := realTime{realTimeConnection}

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

	var ping RealTimePing
	err = json.Unmarshal(event.Payload(), &ping)

	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if ping.EventType != "ping" {
		t.Fatal("Payload type should be `ping`", ping.EventType)
	}

	if ping.ID <= 0 {
		t.Fatal("ID should be > 0")
	}
}

func TestPingError(t *testing.T) {
	realTimeConnection := TestRealTimeConnection{
		send: func(e realtime.Event) error {
			return fmt.Errorf("Ping error")
		},
	}

	rtm := realTime{realTimeConnection}

	err := rtm.Ping()
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestPostMessage(t *testing.T) {
	var event realtime.Event
	realTimeConnection := TestRealTimeConnection{
		send: func(e realtime.Event) error {
			if event != nil {
				t.Fatal("realTimeConnection.Send called more than once")
			}
			event = e
			return nil
		},
	}

	rtm := realTime{realTimeConnection}

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

	var message RealTimeMessage
	err = json.Unmarshal(event.Payload(), &message)

	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if message.EventType != "message" {
		t.Fatal("Payload type should be `message`", message.EventType)
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
		send: func(e realtime.Event) error {
			return fmt.Errorf("PostMessage error")
		},
	}

	rtm := realTime{realTimeConnection}

	channel := "#F00DD00D"
	text := "this is the message"
	err := rtm.PostMessage(channel, text)
	if err == nil {
		t.Fatal("Expected error")
	}
}
