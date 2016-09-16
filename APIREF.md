# Guac API Reference

Install:

```
go get github.com/doozr/guac
```

Use:

```go
import "github.com/doozr/guac"
```

Package guac contains a client for the Slack Web and Real Time APIs in Go.

## Usage

#### type ChannelInfo

```go
type ChannelInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsMember bool   `json:"is_member"`
}
```

ChannelInfo represents a channel or group's info.

#### type MessageEvent

```go
type MessageEvent struct {
	Type    string `json:"type"`
	ID      uint64 `json:"id"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
}
```

MessageEvent is a chat message sent to a user or channel.

#### type PingPongEvent

```go
type PingPongEvent struct {
	Type string `json:"type"`
	ID   uint64 `json:"id"`
}
```

PingPongEvent is a ping and also the reciprocal pong.

#### type RealTimeClient

```go
type RealTimeClient struct {
	WebClient
}
```

RealTimeClient is a client of the Slack RealTime API.

The connection stays open between calls until Close is called. If an error is
returned at any point, it should be considered fatal for the connection and a
new connection should be opened with WebClient.RealTime.

Subsequent calls after an error will result in the same error.

#### func (RealTimeClient) Close

```go
func (g RealTimeClient) Close()
```
Close terminates the connection.

#### func (RealTimeClient) ID

```go
func (g RealTimeClient) ID() string
```
ID of the bot

#### func (RealTimeClient) Name

```go
func (g RealTimeClient) Name() string
```
Name of the bot

#### func (RealTimeClient) Ping

```go
func (g RealTimeClient) Ping() (err error)
```
Ping sends a ping request.

Sends a bare ping with no additional information.

#### func (RealTimeClient) PostMessage

```go
func (g RealTimeClient) PostMessage(channel, text string) (err error)
```
PostMessage sends a chat message to the given channel.

The message is posted as the bot itself, and does not try to take on the
identity of a user. Use the API formatting standard.

#### func (RealTimeClient) Receive

```go
func (g RealTimeClient) Receive() (event interface{}, err error)
```
Receive an event from the Slack RealTime API.

Receive one of the concrete event types and return it. The event should be
checked with a type assertion to determine its type. If a message of an as-yet
unsupported type arrives it will be ignored.

#### type UserChangeEvent

```go
type UserChangeEvent struct {
	Type     string `json:"type"`
	UserInfo `json:"user"`
}
```

UserChangeEvent is a notification that something about a user has changed.
Currently only username changes are supported.

#### type UserInfo

```go
type UserInfo struct {
	ID   string
	Name string
}
```

UserInfo represents a single user's profile info.

#### type WebClient

```go
type WebClient struct {
}
```

WebClient is an interface to the Slack Web API.

#### func  New

```go
func New(token string) WebClient
```
New Slack Web API client.

#### func (WebClient) ChannelsList

```go
func (c WebClient) ChannelsList() (channels []ChannelInfo, err error)
```
ChannelsList gets a list of channel information.

All channels, including archived channels, are returned excluding private
channels. Use GroupsList to retrieve private channels.

#### func (WebClient) GroupsList

```go
func (c WebClient) GroupsList() (channels []ChannelInfo, err error)
```
GroupsList gets a list of private channel information.

All private channels, but not single or multi-user IMs.

#### func (WebClient) IMOpen

```go
func (c WebClient) IMOpen(user string) (channel string, err error)
```
IMOpen opens or returns an IM channel with a specified user.

If an IM with the specified user already exists and is not archived it is
returned, otherwise a new IM channel is opened with that user.

#### func (WebClient) PersistentRealTime

```go
func (c WebClient) PersistentRealTime() (client RealTimeClient, err error)
```
PersistentRealTime connects to the Slack RealTime API using the Web client's
credentials and reconnects whenever the connection drops.

The only way to stop it reconnecting is to use RealTimeClient.Close()

#### func (WebClient) RealTime

```go
func (c WebClient) RealTime() (client RealTimeClient, err error)
```
RealTime connects to the Slack RealTime API using the Web client's credentials.

The returned object represents a websocket connection that remains open between
calls until the Close method is called.

#### func (WebClient) UsersList

```go
func (c WebClient) UsersList() (users []UserInfo, err error)
```
UsersList returns a list of user information.

All users are returned, including deleted and deactivated users.
