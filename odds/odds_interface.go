package odds

type SportsbookClient interface {
	GetSportingLeagues() (leagues []SportingLeague, err error)
	GetUpcomingEvents(sportKey string) (events []SportingEvent, err error)
	GetOddsForEvents(sportKey string, regions []string, markets []string) (odds []SportingOdds, err error)
}
