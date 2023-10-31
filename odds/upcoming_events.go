package odds

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func (c *Client) GetUpcomingEvents(sportKey string) (events []SportingEvent, err error) {

	input := RequestInput{
		Events:   true,
		SportKey: sportKey,
	}

	response, err := c.Request(input)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response.Body(), &events); err != nil {
		log.Err(err).Msg("could not unmarshal response from odds api")
		return nil, err
	}

	return events, nil
}
