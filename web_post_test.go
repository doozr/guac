package guac

import (
	"net/url"
	"testing"

	"github.com/doozr/guac/web"
)

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
