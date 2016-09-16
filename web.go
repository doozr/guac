package guac

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/doozr/guac/persistent"
	"github.com/doozr/guac/realtime"
	"github.com/doozr/guac/web"
)

// WebClient is an interface to the Slack Web API.
type WebClient struct {
	client web.Client
}

// RealTime connects to the Slack RealTime API using the Web client's credentials.
//
// The returned object represents a websocket connection that remains open
// between calls until the Close method is called.
func (c WebClient) RealTime() (client RealTimeClient, err error) {
	websocketConn, err := realtime.New(c.client).Dial()
	if err != nil {
		return
	}

	client = RealTimeClient{
		WebClient:  c,
		connection: websocketConn,
	}
	return
}

// PersistentRealTime connects to the Slack RealTime API using the Web client's credentials
// and reconnects whenever the connection drops.
//
// The only way to stop it reconnecting is to use RealTimeClient.Close().
//
// The timeout parameter is the time after which an open connection is considered
// inactive. If this timeout is hit the client will reconnect.
func (c WebClient) PersistentRealTime(timeout time.Duration) (client RealTimeClient, err error) {
	dialer := realtime.New(c.client)
	websocketConn := persistent.New(dialer, timeout)
	client = RealTimeClient{
		WebClient:  c,
		connection: websocketConn,
	}
	return
}

// UsersList returns a list of user information.
//
// All users are returned, including deleted and deactivated users.
func (c WebClient) UsersList() (users []UserInfo, err error) {
	response, err := c.client.Get("users.list", nil)
	if err != nil {
		return
	}

	if !response.Success() {
		err = response.Error()
		return
	}

	userList := struct {
		Users []UserInfo `json:"users"`
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
func (c WebClient) ChannelsList() (channels []ChannelInfo, err error) {
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
func (c WebClient) GroupsList() (channels []ChannelInfo, err error) {
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
func (c WebClient) IMOpen(user string) (channel string, err error) {
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
