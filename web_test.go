package guac

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/doozr/guac/web"
)

type TestWebClient struct {
	get  func(string, url.Values) (web.Response, error)
	post func(string, url.Values) (web.Response, error)
}

func (c TestWebClient) Get(endPoint string, values url.Values) (web.Response, error) {
	return c.get(endPoint, values)
}

func (c TestWebClient) Post(endPoint string, values url.Values) (web.Response, error) {
	return c.post(endPoint, values)
}

type TestWebResponse struct {
	err     error
	payload []byte
}

func (t TestWebResponse) Success() bool {
	return t.err == nil
}

func (t TestWebResponse) Error() error {
	return t.err
}

func (t TestWebResponse) Payload() []byte {
	return t.payload
}

func checkEndpoint(
	t *testing.T,
	endPoint string,
	payload []byte,
	expected interface{},
	fn func(WebClient) (interface{}, error)) {

	testClient := TestWebClient{
		get: func(ep string, values url.Values) (web.Response, error) {
			if endPoint != ep {
				t.Fatal("Incorrect endpoint called", endPoint, ep)
			}
			return TestWebResponse{
				payload: payload,
			}, nil
		},
	}
	client := WebClient{client: testClient}

	result, err := fn(client)
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatal("Result did not match", expected, result)
	}
}

func TestUsersList(t *testing.T) {
	payload := []byte(`{ "members": [ {
						"id": "U1234",
						"name": "Mr Test"
					}, {
						"id": "U5678",
						"name": "Testy McTesterson"
					}]
				}`)

	expected := []UserInfo{
		UserInfo{
			ID:   "U1234",
			Name: "Mr Test",
		},
		UserInfo{
			ID:   "U5678",
			Name: "Testy McTesterson",
		},
	}

	fn := func(client WebClient) (interface{}, error) {
		return client.UsersList()
	}

	checkEndpoint(t, "users.list", payload, expected, fn)
}

func TestChannelsList(t *testing.T) {
	payload := []byte(`{ "channels": [ {
			"id": "C1234",
			"name": "#general",
			"is_member": true
		}, {
			"id": "C5678",
			"name": "#random"
		}]
	}`)

	expected := []ChannelInfo{
		ChannelInfo{
			ID:       "C1234",
			Name:     "#general",
			IsMember: true,
		},
		ChannelInfo{
			ID:       "C5678",
			Name:     "#random",
			IsMember: false,
		},
	}

	fn := func(client WebClient) (interface{}, error) {
		return client.ChannelsList()
	}

	checkEndpoint(t, "channels.list", payload, expected, fn)
}

func TestGroupsList(t *testing.T) {
	payload := []byte(`{ "groups": [ {
			"id": "C1234",
			"name": "#general",
			"is_member": true
		}, {
			"id": "C5678",
			"name": "#random"
		}]
	}`)

	expected := []ChannelInfo{
		ChannelInfo{
			ID:       "C1234",
			Name:     "#general",
			IsMember: true,
		},
		ChannelInfo{
			ID:       "C5678",
			Name:     "#random",
			IsMember: false,
		},
	}

	fn := func(client WebClient) (interface{}, error) {
		return client.GroupsList()
	}

	checkEndpoint(t, "groups.list", payload, expected, fn)
}

func TestIMOpen(t *testing.T) {
	testClient := TestWebClient{
		post: func(ep string, values url.Values) (web.Response, error) {
			if "im.open" != ep {
				t.Fatal("Incorrect endpoint called", "im.open", ep)
			}
			if values.Get("user") != "U1234" {
				t.Fatal("User does not match", "U1234", values.Get("user"))
			}
			return TestWebResponse{
				payload: []byte(`{
						"channel": {
							"ID": "D123456"
						}
					}`),
			}, nil
		},
	}
	client := WebClient{client: testClient}

	channel, err := client.IMOpen("U1234")
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if channel != "D123456" {
		t.Fatal("Channel does not match", "D123456", channel)
	}
}
