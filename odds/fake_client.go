package odds

import (
	"encoding/json"
)

type FakeClient struct {
}

func NewFakeClient(apiKey string) *FakeClient {
	return &FakeClient{}
}

func (f FakeClient) GetSportingLeagues() (leagues []SportingLeague, err error) {

	if err = json.Unmarshal([]byte(FakeLeaguesResponse), leagues); err != nil {
		return nil, err
	}

	return leagues, nil
}

func (f FakeClient) GetOddsForEvents(sportKey string, regions []string, marketsInput []string) (odds []SportingOdds, err error) {

	if err = json.Unmarshal([]byte(FakeOddsResponse), odds); err != nil {
		return nil, err
	}

	return odds, nil
}

func (f FakeClient) GetUpcomingEvents(sportKey string) (events []SportingEvent, err error) {

	if err = json.Unmarshal([]byte(FakeUpcomingEventsResponse), events); err != nil {
		return nil, err
	}

	return events, nil
}
