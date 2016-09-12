package guac

import "github.com/doozr/guac/web"

// New Slack Web API client
func New(token string) WebClient {
	return WebClient{
		client: web.New(token),
	}
}
