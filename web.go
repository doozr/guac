package guac

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/doozr/guac/realtime"
	"github.com/doozr/guac/web"
)

// Default implementation of WebClient
type webClient struct {
	client web.Client
}

// RealTime connects to the Slack RealTime API using the Web client's
// credentials.
//
// The returned object represents a websocket connection that remains open
// between calls until the Close method is called.
func (c *webClient) RealTime() (client RealTimeClient, err error) {
	websocketConn, err := realtime.New(c.client).Dial()
	if err != nil {
		return
	}

	client = &realTimeClient{
		WebClient:  c,
		connection: websocketConn,
	}
	return
}

// UsersList returns a list of user information.
//
// All users are returned, including deleted and deactivated users.
func (c *webClient) UsersList() (users []UserInfo, err error) {
	response, err := c.client.Get("users.list", nil)
	if err != nil {
		return
	}

	if !response.Success() {
		err = response.Error()
		return
	}

	userList := struct {
		Users []UserInfo `json:"members"`
	}{}

	err = json.Unmarshal(response.Payload(), &userList)
	if err != nil {
		return
	}

	users = userList.Users
	return
}

// ChannelsList gets a list of channel information.
//
// All channels, including archived channels, are returned excluding private
// channels. Use GroupsList to retrieve private channels.
func (c *webClient) ChannelsList() (channels []ChannelInfo, err error) {
	response, err := c.client.Get("channels.list", nil)
	if err != nil {
		return
	}

	if !response.Success() {
		err = response.Error()
		return
	}

	channelList := struct {
		Channels []ChannelInfo `json:"channels"`
	}{}

	err = json.Unmarshal(response.Payload(), &channelList)
	if err != nil {
		return
	}

	channels = channelList.Channels
	return
}

// GroupsList gets a list of private channel information.
//
// All private channels, but not single or multi-user IMs.
func (c *webClient) GroupsList() (channels []ChannelInfo, err error) {
	response, err := c.client.Get("groups.list", nil)
	if err != nil {
		return
	}

	if !response.Success() {
		err = response.Error()
		return
	}

	groupList := struct {
		Groups []ChannelInfo `json:"groups"`
	}{}

	err = json.Unmarshal(response.Payload(), &groupList)
	if err != nil {
		return
	}

	channels = groupList.Groups
	return
}

// IMOpen opens or returns an IM channel with a specified user.
//
// If an IM with the specified user already exists and is not archived it is
// returned, otherwise a new IM channel is opened with that user.
func (c *webClient) IMOpen(user string) (channel string, err error) {
	values := url.Values{}
	values.Add("user", user)
	response, err := c.client.Post("im.open", values)
	if err != nil {
		return
	}

	if !response.Success() {
		err = response.Error()
		return
	}

	result := struct {
		Channel struct {
			ID string `json:"id"`
		} `json:"channel"`
	}{}
	err = json.Unmarshal(response.Payload(), &result)
	if err != nil {
		return
	}

	channel = result.Channel.ID
	if channel == "" {
		err = fmt.Errorf("Could not retrieve channel ID")
	}
	return
}

func (c *webClient) PostMessage(channel, text string) (err error) {
	values := url.Values{}
	values.Add("channel", channel)
	values.Add("text", text)
	response, err := c.client.Post("chat.postMessage", values)
	if err != nil {
		return
	}

	if !response.Success() {
		err = response.Error()
	}

	return
}

func (c *webClient) PostSnippet(channel, content, filename, filetype, title, initialComment string) (err error) {
	values := url.Values{}
	values.Add("channel", channel)
	values.Add("content", content)

	if filename != "" {
		values.Add("filename", filename)
	}

	if filetype != "" {
		values.Add("filetype", filetype)
	}

	if initialComment != "" {
		values.Add("initial_comment", initialComment)
	}

	response, err := c.client.Post("files.upload", values)
	if err != nil {
		return
	}

	if !response.Success() {
		err = response.Error()
	}

	return
}
