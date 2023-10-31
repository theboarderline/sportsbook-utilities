package odds

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"strings"
)

type RequestInput struct {
	Events   bool
	SportKey string
	Regions  []string
	Markets  []string
}

func (c *Client) Request(input RequestInput) (response *resty.Response, err error) {

	response, err = c.httpClient.R().
		SetHeader(RapidAPIKeyHeaderName, c.apiKey).
		SetHeader(RapidAPIHostHeaderName, RapidAPIHostHeaderValue).
		SetQueryParams(input.GetQueryParams()).
		Get(input.FormatURL())

	if err != nil {
		log.Err(err).Msgf("could get data from odds api: %s", input.FormatURL())
		return nil, err
	}

	if response.IsError() {
		err = errors.New(string(response.Body()))
		log.Err(err).Msgf("odds response status: %d", response.StatusCode())
		return nil, err
	}

	log.Debug().Msg(string(response.Body()))
	return response, nil
}

func (r RequestInput) FormatURL() string {
	baseURL := fmt.Sprintf("%s/sports", RapidAPIEndpoint)

	if r.SportKey != "" {
		baseURL = fmt.Sprintf("%s/%s", baseURL, r.SportKey)
	}

	if r.Events {
		baseURL = fmt.Sprintf("%s/scores", baseURL)

	} else if len(r.Markets) > 0 {
		baseURL = fmt.Sprintf("%s/odds", baseURL)
	}

	return baseURL
}

func (r RequestInput) GetQueryParams() map[string]string {
	params := map[string]string{
		"markets": r.markets(),
		"regions": r.regions(),
	}

	return params
}

func (r RequestInput) markets() string {

	return strings.Join(r.Markets, ",")
}

func (r RequestInput) regions() string {
	return strings.Join(r.Regions, ",")
}
