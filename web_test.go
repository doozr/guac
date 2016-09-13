package guac

import (
	"net/url"

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
