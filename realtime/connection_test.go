package realtime

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type TestRawConnection struct {
	Incoming chan []byte
	Outgoing chan []byte
}

func (c TestRawConnection) ID() string {
	return "testing"
}

func (c TestRawConnection) Receive() (payload []byte, err error) {
	v := <-c.Incoming
	if v == nil {
		return nil, fmt.Errorf("Incoming Error")
	}
	return v, nil
}

func (c TestRawConnection) Send(payload []byte) (err error) {
	if payload == nil {
		return fmt.Errorf("Outgoing Error")
	}
	c.Outgoing <- payload
	return
}

func NewTestRawConnection(inBuffer, outBuffer int) TestRawConnection {
	return TestRawConnection{
		Incoming: make(chan []byte, inBuffer),
		Outgoing: make(chan []byte, outBuffer),
	}
}

func TestSuccessfulReceive(t *testing.T) {
	raw := NewTestRawConnection(1, 0)
	conn := New(raw)

	raw.Incoming <- []byte(`{
			"type": "test",
			"body": "this is a test"
		}`)

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
	raw := NewTestRawConnection(1, 0)
	conn := New(raw)

	raw.Incoming <- []byte(`{
			"type": "test",
			INVALID JSON
		}`)

	result, err := conn.Receive()
	if result != nil {
		t.Fatal("Expected nil result", result)
	}

	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestReceiveError(t *testing.T) {
	raw := NewTestRawConnection(0, 0)
	conn := New(raw)

	close(raw.Incoming)

	result, err := conn.Receive()
	if result != nil {
		t.Fatal("Expected nil result", result)
	}

	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestSuccessfulSend(t *testing.T) {
	raw := NewTestRawConnection(0, 1)
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

	payload := <-raw.Outgoing

	if !reflect.DeepEqual(payload, event.Payload()) {
		t.Fatal("Payload did not match", event.Payload(), payload)
	}
}

func TestSendError(t *testing.T) {
	raw := NewTestRawConnection(0, 1)
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
