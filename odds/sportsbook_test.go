package odds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/sportsbook-utilities/odds"
	"os"
)

var _ = Describe("Sportsbook", func() {

	var (
		sportsbookClient *odds.Client
	)

	BeforeEach(func() {
		sportsbookClient = odds.NewClient(os.Getenv("RAPID_API_KEY"))
	})

	It("can format the url for the sportsbook", func() {
		input := odds.RequestInput{
			SportKey: "americanfootball_ncaaf",
			Markets:  []string{"h2h", "spreads"},
		}
		url := input.FormatURL()
		Expect(url).To(Equal("https://odds.p.rapidapi.com/v4/sports/americanfootball_ncaaf/odds"))
	})

	It("can get the query parameters for the sportsbook", func() {
		input := odds.RequestInput{
			Markets: []string{"h2h", "spreads"},
			Regions: []string{"us,uk"},
		}
		params := input.GetQueryParams()
		Expect(params).To(HaveKeyWithValue("markets", "h2h,spreads"))
		Expect(params).To(HaveKeyWithValue("regions", "us,uk"))
		Expect(params).To(HaveKeyWithValue("oddsFormat", "american"))
	})

	It("can get the upcoming sporting leagues from the odds", func() {
		leagues, err := sportsbookClient.GetSportingLeagues()
		Expect(err).NotTo(HaveOccurred())
		Expect(leagues).NotTo(BeEmpty())
		Expect(len(leagues)).To(BeNumerically(">=", 0))
		Expect(leagues[0].SportKey).NotTo(BeEmpty())
	})

	It("can get the upcoming games for a sport from the odds", func() {
		events, err := sportsbookClient.GetUpcomingEvents("americanfootball_ncaaf")
		Expect(err).NotTo(HaveOccurred())
		Expect(events).NotTo(BeEmpty())
		Expect(len(events)).To(BeNumerically(">=", 0))
	})

	It("can get the odds for a game from the odds", func() {
		odds, err := sportsbookClient.GetOddsForEvents("americanfootball_ncaaf", []string{"us"}, []string{"h2h"})
		Expect(err).NotTo(HaveOccurred())
		Expect(odds).NotTo(BeEmpty())
		Expect(len(odds)).To(BeNumerically(">=", 0))
		Expect(len(odds[0].Bookmakers)).To(BeNumerically(">=", 0))
		Expect(odds[0].Bookmakers[0].Key).NotTo(BeEmpty())
	})

})
