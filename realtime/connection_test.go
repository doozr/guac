package realtime

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type TestRawConnection struct {
	closed  chan bool
	receive func() ([]byte, error)
	send    func([]byte) error
}

func (c TestRawConnection) Close() {
	close(c.closed)
}

func (c TestRawConnection) ID() string {
	return "testing"
}

func (c TestRawConnection) Receive() (payload []byte, err error) {
	return c.receive()
}

func (c TestRawConnection) Send(payload []byte) (err error) {
	return c.send(payload)
}

func TestClose(t *testing.T) {
	raw := TestRawConnection{
		closed: make(chan bool),
	}
	conn := New(raw)

	conn.Close()

	select {
	case <-raw.closed: // Responds if channel is closed
	default:
		t.Fatal("Expected connection to be closed")
	}
}

func TestSuccessfulReceive(t *testing.T) {
	called := false
	raw := TestRawConnection{
		receive: func() ([]byte, error) {
			if called {
				t.Fatal("RawConnection.Receive called more than once")
			}
			called = true
			return []byte(`{ "type": "test", "body": "this is a test" }`), nil
		},
	}
	conn := New(raw)

	result, err := conn.Receive()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if result.EventType() != "test" {
		t.Fatal("Event types did not match", "test", result.EventType())
	}

	v := make(map[string]string)
	err = json.Unmarshal(result.Payload(), &v)
	if err != nil {
		t.Fatal("Failed to unmarshal payload", err)
	}

	expected := map[string]string{
		"type": "test",
		"body": "this is a test",
	}

	if !reflect.DeepEqual(expected, v) {
		t.Fatal("Payload did not match", expected, v)
	}
}

func TestInvalidJson(t *testing.T) {
	called := false
	raw := TestRawConnection{
		receive: func() ([]byte, error) {
			if called {
				t.Fatal("RawConnection.Receive called more than once")
			}
			called = true
			return []byte(`{ "type": "test", INVALID JSON }`), nil
		},
	}
	conn := New(raw)

	result, err := conn.Receive()
	if result != nil {
		t.Fatal("Expected nil result", result)
	}

	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestReceiveError(t *testing.T) {
	raw := TestRawConnection{
		receive: func() ([]byte, error) {
			return nil, fmt.Errorf("Receive error")
		},
	}
	conn := New(raw)

	result, err := conn.Receive()
	if result != nil {
		t.Fatal("Expected nil result", result)
	}

	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestSuccessfulSend(t *testing.T) {
	var payload []byte
	raw := TestRawConnection{
		send: func(p []byte) error {
			if payload != nil {
				t.Fatal("RawConnection.Send called more than once")
			}
			payload = p
			return nil
		},
	}
	conn := New(raw)

	event := realTimeEvent{
		eventType: "test",
		payload: []byte(`{
				"type": "test",
				"body": "this is a test"
			}`),
	}

	err := conn.Send(event)
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if !reflect.DeepEqual(payload, event.Payload()) {
		t.Fatal("Payload did not match", event.Payload(), payload)
	}
}

func TestSendError(t *testing.T) {
	raw := TestRawConnection{
		send: func(p []byte) error {
			return fmt.Errorf("Outgoing error")
		},
	}
	conn := New(raw)

	event := realTimeEvent{
		eventType: "test",
		payload:   nil,
	}

	err := conn.Send(event)
	if err == nil {
		t.Fatal("Expected error")
	}
}
