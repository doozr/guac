# guac

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
slack := guac.New(token)
```

Connecting to the Real Time API is done via an existing web client and opens a
websocket to communicate with the Slack service.

```go
rtm := slack.RealTime()
```

Receive events via the `Receive()` method. All events are returned from the same
function so the best way to handle them is with a type switch. this could call
handlers, push the events onto channels, or anything else.

```go
func receiveEvents(rtm slack.RealTimeClient,
                   done chan struct{},
                   messages chan MessageEvent,
                   userChanges chan UserChangeEvent) {
    for {
        select {
        case <-done:
            return
        default:
            e, err := rtm.Receive()
            if err != nil {
                // Cannot continue with this instance
                return
            }

            switch event := e.(type) {
            case MessageEvent:
                messages <- event
            case UserChangeEvent:
                userChanges <- event
            default:
                // Unhandled
        }
    }
}
```

To force the `Receive()` method to stop listening, close the connection with the
`Close()` method. This will force the `Receive()` method to return an error.
Once an error is received from the Receive method, it is considered terminal.
Further calls will return the same error. At this point, create a new connection
with `WebClient.RealTime()`.

## Limitations

Only a very small subset of the Slack API functionality is implemented, although
the code is extensible enough to allow quick addition of new endpoints and event
types. It should be enough to get started with a very basic conversational bot.
