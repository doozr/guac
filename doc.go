/*Package guac provides clients to connect bots to Slack.

Getting Started

Instantiating the client requires only a bot integration token,
configured via the *Custom Integrations* section of the Slack admin
panel. Bots can discover their own name and channels via the API
itself so none of that information is required.

	slack := guac.New(token)

Connecting to the Real Time API is done via an existing web client
and opens a websocket to communicate with the Slack service.

	rtm := slack.RealTime()

*/
package guac
