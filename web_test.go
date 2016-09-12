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

func TestUsersList(t *testing.T) {
	testClient := TestWebClient{
		get: func(endPoint string, values url.Values) (web.Response, error) {
			return TestWebResponse{
				payload: []byte(`{ "users": [ {
						"id": "U1234",
						"name": "Mr Test"
					}, {
						"id": "U5678",
						"name": "Testy McTesterson"
					}]
				}`),
			}, nil
		},
	}
	client := WebClient{client: testClient}

	users, err := client.UsersList()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

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
	if !reflect.DeepEqual(expected, users) {
		t.Fatal("Result did not match", expected, users)
	}
}
