package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/doozr/jot"
)

// Client of the Web API.
type client struct {
	token string
}

// New web API client.
func New(token string) Client {
	return client{
		token: token,
	}
}

// Get an event from the API.
func (c client) Get(endPoint string, values url.Values) (response Response, err error) {
	if values == nil {
		values = url.Values{}
	}

	values.Add("token", c.token)
	url := fmt.Sprintf("https://slack.com/api/%s?%s", endPoint, values.Encode())

	jot.Print("Web GET: ", url)
	resp, err := http.Get(url)
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
	jot.Printf("Web GET to %s received %s", url, string(bytes))

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
	jot.Printf("Web POST: %s with form values: %s", url, values.Encode())
	resp, err := http.PostForm(url, values)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("API POST '%s' failed with code %d", endPoint, resp.StatusCode)
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	jot.Printf("Web POST to %s received %s", url, string(bytes))

	respObj := apiResponse{}
	err = json.Unmarshal(bytes, &respObj)
	respObj.Raw = bytes
	response = respObj

	return
}
