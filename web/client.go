package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/doozr/jot"
)

// httpClient interface we will use to actually do requests
type httpClient interface {
	Get(string) (*http.Response, error)
	PostForm(string, url.Values) (*http.Response, error)
}

// Client of the Web API.
type client struct {
	token string
	http  httpClient
}

// New web API client.
func New(token string, http httpClient) Client {
	return client{
		token: token,
		http:  http,
	}
}

// Get an event from the API.
func (c client) Get(endPoint string, values url.Values) (response Response, err error) {
	if values == nil {
		values = url.Values{}
	}

	values.Add("token", c.token)
	url := fmt.Sprintf("https://slack.com/api/%s?%s", endPoint, values.Encode())

	jot.Print("web.client: GET ", url)
	resp, err := c.http.Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("API GET '%s' failed with code %d", url, resp.StatusCode)
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	if jot.Enabled() {
		o := string(bytes)
		if len(o) > 256 {
			o = o[:253] + "..."
		}
		jot.Printf("web.client: GET %s received %s", url, o)
	}

	respObj := apiResponse{}
	err = json.Unmarshal(bytes, &respObj)
	respObj.Raw = bytes
	response = respObj

	return
}

// Post an action request and return the response.
func (c client) Post(endPoint string, values url.Values) (response Response, err error) {
	if values == nil {
		values = url.Values{}
	}

	values.Add("token", c.token)

	url := "https://slack.com/api/" + endPoint
	jot.Printf("web.client: POST to %s with form values: %s", url, values.Encode())
	resp, err := c.http.PostForm(url, values)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("POST to %s failed with code %d", endPoint, resp.StatusCode)
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	if jot.Enabled() {
		o := string(bytes)
		if len(o) > 256 {
			o = o[:253] + "..."
		}
		jot.Printf("web.client: POST %s received %s", url, o)
	}

	respObj := apiResponse{}
	err = json.Unmarshal(bytes, &respObj)
	respObj.Raw = bytes
	response = respObj

	return
}
