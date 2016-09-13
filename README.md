# guac

Install:

    go get github.com/doozr/guac

And use:

    import "github.com/doozr/guac"

Package guac provides clients to connect bots to Slack.


### Getting Started

Instantiating the client requires only a bot integration token, configured via
the *Custom Integrations* section of the Slack admin panel. Bots can discover
their own name and channels via the API itself so none of that information is
required.

    slack := guac.New(token)

Connecting to the Real Time API is done via an existing web client and opens a
websocket to communicate with the Slack service.

    rtm := slack.RealTime()

## Usage

#### type ChannelInfo

```go
type ChannelInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsMember bool   `json:"is_member"`
}
```

ChannelInfo represents a channel or group's info

#### type RealTimeClient

```go
type RealTimeClient struct {
}
```

RealTimeClient is a client of the Slack RealTime API

#### func (RealTimeClient) Ping

```go
func (g RealTimeClient) Ping() (err error)
```
Ping sends a ping request

#### func (RealTimeClient) PostMessage

```go
func (g RealTimeClient) PostMessage(channel, text string) (err error)
```
PostMessage sends a chat message to the given channel

#### func (RealTimeClient) Receive

```go
func (g RealTimeClient) Receive() (event interface{}, err error)
```
Receive an event from the Slack RealTime API

#### type RealTimeMessage

```go
type RealTimeMessage struct {
	Type    string `json:"type"`
	ID      uint64 `json:"id"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
}
```

RealTimeMessage is a chat message sent to a user or channel

#### type RealTimePing

```go
type RealTimePing struct {
	Type string `json:"type"`
	ID   uint64 `json:"id"`
}
```

RealTimePing is a ping and also the reciprocal pong

#### type RealTimeUserChange

```go
type RealTimeUserChange struct {
	Type string   `json:"type"`
	User UserInfo `json:"user"`
}
```

RealTimeUserChange is a notification that something about a user has changed
Currently only username changes are supported

#### type UserInfo

```go
type UserInfo struct {
	ID   string
	Name string
}
```

UserInfo represents a single user's profile info

#### type WebClient

```go
type WebClient struct {
}
```

WebClient is an interface to the Slack Web API

#### func  New

```go
func New(token string) WebClient
```
New Slack Web API client

#### func (WebClient) ChannelsList

```go
func (c WebClient) ChannelsList() (channels []ChannelInfo, err error)
```
ChannelsList gets a list of channel information

#### func (WebClient) GroupsList

```go
func (c WebClient) GroupsList() (channels []ChannelInfo, err error)
```
GroupsList gets a list of private channel information Slack's nomenclature for
different types of channel is weird

#### func (WebClient) IMOpen

```go
func (c WebClient) IMOpen(user string) (channel string, err error)
```
IMOpen opens or returns an IM channel with a specified user

#### func (WebClient) RealTime

```go
func (c WebClient) RealTime() (client RealTimeClient, err error)
```
RealTime connects to the Slack RealTime API using the Web client's credentials

#### func (WebClient) UsersList

```go
func (c WebClient) UsersList() (users []UserInfo, err error)
```
UsersList returns a list of user information
