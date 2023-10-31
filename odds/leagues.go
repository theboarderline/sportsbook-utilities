package odds

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func (c *Client) GetSportingLeagues() (leagues []SportingLeague, err error) {

	res, err := c.Request(RequestInput{})
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res.Body(), &leagues); err != nil {
		log.Err(err).Msg("could not unmarshal response from odds api")
		return nil, err
	}

	return leagues, nil
}
