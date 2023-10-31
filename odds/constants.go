package odds

import "fmt"

const (
	RapidAPIKeyHeaderName   = "X-RapidAPI-Key"
	RapidAPIHostHeaderName  = "X-RapidAPI-Host"
	RapidAPIHostHeaderValue = "odds.p.rapidapi.com"
)

var (
	RapidAPIEndpoint = fmt.Sprintf("https://%s/v4", RapidAPIHostHeaderValue)
)
