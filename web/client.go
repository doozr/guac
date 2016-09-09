package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/doozr/guac/slack"
)

// Payload of the response
func (t apiResponse) Payload() []byte {
	return t.Raw
}

// Client of the Web API
type Client struct {
	token string
}

// New web API client
func New(token string) Client {
	return Client{
		token: token,
	}
}

// Get an event from the API
func (c Client) Get(endPoint string, values url.Values) (response slack.APIResponse, err error) {
	if values == nil {
		values = url.Values{}
	}

	values.Add("token", c.token)
	url := fmt.Sprintf("https://slack.com/api/%s?%s", endPoint, values.Encode())
	log.Printf("Web API GET %s", url)

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

	respObj := apiResponse{}
	err = json.Unmarshal(bytes, &respObj)
	respObj.Raw = bytes
	response = respObj

	return
}

// Post an action request and return the response
func (c Client) Post(endPoint string, values url.Values) (response slack.APIResponse, err error) {
	if values == nil {
		values = url.Values{}
	}

	values.Add("token", c.token)
	log.Printf("Web API POST %s with values %s", endPoint, values.Encode())

	resp, err := http.PostForm(endPoint, values)
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

	respObj := apiResponse{}
	err = json.Unmarshal(bytes, &respObj)
	respObj.Raw = bytes
	response = respObj

	return
}

// apiResponse is a concrete implementation of slack.APIResponse
type apiResponse struct {
	OK  bool   `json:"ok"`
	Err string `json:"error"`
	Raw []byte
}

// Success returns true if no error occured
func (t apiResponse) Success() bool {
	return t.OK
}

// Error returns an error containing details of the fault, if there is one
func (t apiResponse) Error() (err error) {
	if !t.Success() {
		err = fmt.Errorf(t.Err)
	}
	return
}
