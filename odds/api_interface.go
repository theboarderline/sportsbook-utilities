package odds

import "time"

type SportingLeague struct {
	SportKey     string `json:"key"`
	Group        string `json:"group"`
	SportTitle   string `json:"title"`
	Description  string `json:"description"`
	Active       bool   `json:"active"`
	HasOutrights bool   `json:"has_outrights"`
}

type SportingEvent struct {
	ID           string    `json:"id"`
	SportKey     string    `json:"sport_key"`
	SportTitle   string    `json:"sport_title"`
	CommenceTime time.Time `json:"commence_time"`
	Completed    bool      `json:"completed"`
	HomeTeam     string    `json:"home_team"`
	AwayTeam     string    `json:"away_team"`
	Scores       string    `json:"scores"`
	LastUpdate   string    `json:"last_update"`
}

type SportingOdds struct {
	SportingEvent
	Bookmakers []Bookmaker `json:"bookmakers"`
}

type Bookmaker struct {
	Key        string    `json:"key"`
	Title      string    `json:"title"`
	LastUpdate time.Time `json:"last_update"`
	Markets    []Market  `json:"markets"`
}

type Market struct {
	Key      string    `json:"key"`
	Outcomes []Outcome `json:"outcomes"`
}

type Outcome struct {
	Label  string  `json:"label"`
	Price  float32 `json:"price"`
	Points float32 `json:"points"`
}

type QueryParams struct {
	Regions string `json:"regions"`
	Markets string `json:"markets"`
}
