# Guac

<img align="right" width="180" style="margin: 12px" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4NCjwhRE9DVFlQRSBzdmcgUFVC%0D%0ATElDICItLy9XM0MvL0RURCBTVkcgMS4xLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL0dyYXBoaWNz%0D%0AL1NWRy8xLjEvRFREL3N2ZzExLmR0ZCI+DQo8c3ZnIHZlcnNpb249IjEuMiIgd2lkdGg9IjM4LjYy%0D%0AbW0iIGhlaWdodD0iNjQuMDJtbSIgdmlld0JveD0iNzMyNSA5ODY0IDM4NjIgNDQwMiIgcHJlc2Vy%0D%0AdmVBc3BlY3RSYXRpbz0ieE1pZFlNaWQiIGZpbGwtcnVsZT0iZXZlbm9kZCIgc3Ryb2tlLXdpZHRo%0D%0APSIyOC4yMjIiIHN0cm9rZS1saW5lam9pbj0icm91bmQiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9y%0D%0AZy8yMDAwL3N2ZyIgeG1sbnM6b29vPSJodHRwOi8veG1sLm9wZW5vZmZpY2Uub3JnL3N2Zy9leHBv%0D%0AcnQiIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4bWxuczpwcmVz%0D%0AZW50YXRpb249Imh0dHA6Ly9zdW4uY29tL3htbG5zL3N0YXJvZmZpY2UvcHJlc2VudGF0aW9uIiB4%0D%0AbWxuczpzbWlsPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1NNSUwyMC8iIHhtbG5zOmFuaW09InVy%0D%0AbjpvYXNpczpuYW1lczp0YzpvcGVuZG9jdW1lbnQ6eG1sbnM6YW5pbWF0aW9uOjEuMCIgeG1sOnNw%0D%0AYWNlPSJwcmVzZXJ2ZSI+CiAgPHBhdGggZmlsbD0icmdiKDE3NCwyMDcsMCkiIHN0cm9rZT0ibm9u%0D%0AZSIgZD0iTSA5MjU1LDk4OTAgQyAxMDMzNSw5ODkwIDExMTYwLDExMjY1IDExMTYwLDEzMDY1IDEx%0D%0AMTYwLDE0ODY1IDk3NjMsMTUwOTcgOTI1NSwxNTA5NyA4NzQ3LDE1MDk3IDczNTAsMTQ4NjUgNzM1%0D%0AMCwxMzA2NSA3MzUwLDExMjY1IDgxNzUsOTg5MCA5MjU1LDk4OTAgWiBNIDczNTAsOTg5MCBMIDcz%0D%0ANTAsOTg5MCBaIE0gMTExNjEsMTYyNDEgTCAxMTE2MSwxNjI0MSBaIi8+CiAgPHBhdGggZmlsbD0i%0D%0Abm9uZSIgc3Ryb2tlPSJyZ2IoODcsMTU3LDI4KSIgc3Ryb2tlLXdpZHRoPSI1MSIgc3Ryb2tlLWxp%0D%0AbmVqb2luPSJyb3VuZCIgZD0iTSA5MjU1LDk4OTAgQyAxMDMzNSw5ODkwIDExMTYwLDExMjY1IDEx%0D%0AMTYwLDEzMDY1IDExMTYwLDE0ODY1IDk3NjMsMTUwOTcgOTI1NSwxNTA5NyA4NzQ3LDE1MDk3IDcz%0D%0ANTAsMTQ4NjUgNzM1MCwxMzA2NSA3MzUwLDExMjY1IDgxNzUsOTg5MCA5MjU1LDk4OTAgWiIvPg0K%0D%0AICA8cGF0aCBmaWxsPSJyZ2IoMjU1LDE0OSwxNCkiIHN0cm9rZT0ibm9uZSIgZD0iTSA5MTkwLDEx%0D%0AOTIyIEMgOTY1OCwxMTkyMiAxMDAxNiwxMjQ3MiAxMDAxNiwxMzE5MiAxMDAxNiwxMzkxMiA5NjU4%0D%0ALDE0NDYyIDkxOTAsMTQ0NjIgODcyMiwxNDQ2MiA4MzY1LDEzOTEyIDgzNjUsMTMxOTIgODM2NSwx%0D%0AMjQ3MiA4NzIyLDExOTIyIDkxOTAsMTE5MjIgWiBNIDgzNjUsMTE5MjIgTCA4MzY1LDExOTIyIFog%0D%0ATSAxMDAxNywxNDQ2MyBMIDEwMDE3LDE0NDYzIFoiLz4KICA8cGF0aCBmaWxsPSJub25lIiBzdHJv%0D%0Aa2U9InJnYigxMjYsMCwzMykiIHN0cm9rZS13aWR0aD0iNTEiIHN0cm9rZS1saW5lam9pbj0icm91%0D%0AbmQiIGQ9Ik0gOTE5MCwxMTkyMiBDIDk2NTgsMTE5MjIgMTAwMTYsMTI0NzIgMTAwMTYsMTMxOTIg%0D%0AMTAwMTYsMTM5MTIgOTY1OCwxNDQ2MiA5MTkwLDE0NDYyIDg3MjIsMTQ0NjIgODM2NSwxMzkxMiA4%0D%0AMzY1LDEzMTkyIDgzNjUsMTI0NzIgODcyMiwxMTkyMiA5MTkwLDExOTIyIFoiLz4NCjwvc3ZnPg0K" />
Guac is a client library for the Slack Web and Real Time APIs in Go. Use it for
writing bots and other integrations. It's not complete but allows basic
conversation via the real time API.

The name comes from a phonetically mangled portmanteau of "Go" and "Slack"
combined with my love of Mexican food.

```
go get github.com/doozr/guac
```

For details of the Guac API see the [API reference](APIREF.md).

## Getting Started

Instantiating the client requires only a bot integration token, configured via
the *Custom Integrations* section of the Slack admin panel. Bots can discover
their own name and channels via the API itself so none of that information is
required.

```go
webClient := guac.New(token)
```

Connecting to the Real Time API is done via an existing client instance, and
opens a websocket to communicate with the Slack service.

```go
rtm, err := webClient.RealTime()
```

Close connections when they're done to clean up any goroutines that are handling
incoming messages from the Slack websocket.

```go
rtm.Close()
```

## Receiving Real Time Events

Receive events via the `Receive()` method. All events are returned from the same
function so the best way to handle them is with a type switch. this could call
handlers, push the events onto channels, or anything else.

```go
func receiveEvents(rtm slack.RealTimeClient,
                   done chan struct{},
                   messages chan guac.MessageEvent,
                   userChanges chan guac.UserChangeEvent) {
    defer func() {
        rtm.Close()
    }()
    for {
        select {
        case <-done:
            return
        default:
            e, err := rtm.Receive()
            if err != nil {
                // log the error
                return
            }
            switch event := e.(type) {
            case guac.MessageEvent:
                messages <- event
            case guac.UserChangeEvent:
                userChanges <- event
            default:
                // Unhandled
        }
    }
}
```

Close the connection with the `Close()` method to stop the `Receive()` method
listening. If an error is received from `Receive()` method it is  considered
terminal. At this point, create a new connection or restart the program.

## Ping Pong

It is recommended to send a Ping request via the `RealTimeClient.Ping()` method
periodically to let Slack know you are still there. It will result in a pong
response returned by the `Receive()` method.

This means that, at a minimum, you should expect a message from Slack at least
as frequently as your pings. If there is no incoming message for a significant
period, it may be that the connection has hung and should be reconnected.

The PersistentRealTime connection accepts a duration to be considered "inactive"
and will reconnect if that timeout is exceeded between messages. Note that it
does not send ping requests; it only times out if nothing comes back. Make sure
pings are being sent to prevent timeout.

## Limitations

Only a very small subset of the Slack API functionality is implemented, although
the code is extensible enough to allow quick addition of new endpoints and event
types. It should be enough to get started with a very basic conversational bot.
