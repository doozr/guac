// Package guac contains a client for the Slack Web and Real Time APIs in Go.
package guac

import (
	"net/http"

	"github.com/doozr/guac/web"
)

// New Slack Web API client.
func New(token string) WebClient {
	return WebClient{
		client: web.New(token, &http.Client{}),
	}
}
