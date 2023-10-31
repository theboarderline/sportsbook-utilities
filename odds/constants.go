package odds

import "fmt"

const (
	RapidAPIKeyHeaderName   = "X-RapidAPI-Key"
	RapidAPIHostHeaderName  = "X-RapidAPI-Host"
	RapidAPIHostHeaderValue = "odds.p.rapidapi.com"
	UpcomingSports          = "upcoming"
	USARegions              = "us"
	H2H                     = "h2h"
	Spreads                 = "spreads"
)

var (
	RapidAPIEndpoint = fmt.Sprintf("https://%s/v4", RapidAPIHostHeaderValue)
)
