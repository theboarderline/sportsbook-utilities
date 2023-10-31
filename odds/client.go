package odds

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	httpClient *resty.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: resty.New(),
		apiKey:     apiKey,
	}
}
