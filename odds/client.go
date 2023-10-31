package odds

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type Client struct {
	httpClient *resty.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: resty.New(),
	}
}

func (c *Client) GetSportingLeagues() (leagues []SportingLeague, err error) {

	url := fmt.Sprintf("%s/sports?all=true", RapidAPIEndpoint)

	response, err := c.httpClient.R().
		SetHeader(RapidAPIKeyHeaderName, os.Getenv("RAPID_API_KEY")).
		SetHeader(RapidAPIHostHeaderName, RapidAPIHostHeaderValue).
		Get(url)

	if err != nil {
		log.Err(err).Msg("could not upload picture to groupme api")
		return nil, err
	}

	if response.IsError() {
		err = errors.New(string(response.Body()))
		log.Err(err).Msgf("status: %d", response.StatusCode())
		return nil, err
	}

	log.Debug().Msg(string(response.Body()))

	if err = json.Unmarshal(response.Body(), &leagues); err != nil {
		log.Err(err).Msg("could not unmarshal response from odds api")
		return nil, err
	}

	return leagues, nil
}

func (c *Client) GetUpcomingEvents(sport string, daysFromInput ...uint) (events []SportingEvent, err error) {

	url := fmt.Sprintf("%s/sports/%s/scores", RapidAPIEndpoint, sport)

	if len(daysFromInput) > 0 {
		url = fmt.Sprintf("%s?daysFrom=%d", url, daysFromInput[0])
	}

	response, err := c.httpClient.R().
		SetHeader(RapidAPIKeyHeaderName, os.Getenv("RAPID_API_KEY")).
		SetHeader(RapidAPIHostHeaderName, RapidAPIHostHeaderValue).
		Get(url)

	if err != nil {
		log.Err(err).Msg("could not upload picture to groupme api")
		return nil, err
	}

	if response.IsError() {
		err = errors.New(string(response.Body()))
		log.Err(err).Msgf("status: %d", response.StatusCode())
		return nil, err
	}

	log.Debug().Msg(string(response.Body()))

	if err = json.Unmarshal(response.Body(), &events); err != nil {
		log.Err(err).Msg("could not unmarshal response from odds api")
		return nil, err
	}

	return events, nil
}

func (c *Client) GetOddsForEvents(sport, region string, marketsInput ...string) (odds []SportingOdds, err error) {

	if sport == "" {
		sport = "upcoming"
	}

	if region == "" {
		region = "us"
	}

	url := fmt.Sprintf("%s/sports/%s/odds?regions=%s", RapidAPIEndpoint, sport, region)

	if len(marketsInput) != 0 {
		markets := strings.Join(marketsInput, ",")
		url = fmt.Sprintf("%s?markets=%s", url, markets)
	}

	response, err := c.httpClient.R().
		SetHeader(RapidAPIKeyHeaderName, os.Getenv("RAPID_API_KEY")).
		SetHeader(RapidAPIHostHeaderName, RapidAPIHostHeaderValue).
		Get(url)

	if err != nil {
		log.Err(err).Msg("could not upload picture to groupme api")
		return nil, err
	}

	if response.IsError() {
		err = errors.New(string(response.Body()))
		log.Err(err).Msgf("status: %d", response.StatusCode())
		return nil, err
	}

	log.Debug().Msg(string(response.Body()))

	if err = json.Unmarshal(response.Body(), &odds); err != nil {
		log.Err(err).Msg("could not unmarshal response from odds api")
		return nil, err
	}

	return odds, nil
}
