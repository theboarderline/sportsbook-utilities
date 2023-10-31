package odds

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func (c *Client) GetOddsForEvents(sportKey string, regions []string, markets []string) (odds []SportingOdds, err error) {

	input := RequestInput{
		SportKey: sportKey,
		Regions:  regions,
		Markets:  markets,
	}

	response, err := c.Request(input)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response.Body(), &odds); err != nil {
		log.Err(err).Msg("could not unmarshal response from odds api")
		return nil, err
	}

	return odds, nil
}
