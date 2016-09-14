/*Package guac provides clients to connect bots to Slack.

Getting Started

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

Receive events via the `RealTime.Receive` method. All events are returned from
the same function so the best way to handle them is with a type switch. this
could call handlers, push the events onto channels, or anything else.

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
                return err
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

*/
package guac
