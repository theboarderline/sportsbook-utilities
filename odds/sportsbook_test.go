package odds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/sportsbook-utilities/odds"
)

var _ = Describe("Sportsbook", func() {

	var (
		sportsbookClient *odds.Client
	)

	BeforeEach(func() {
		sportsbookClient = odds.NewClient()
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
		odds, err := sportsbookClient.GetOddsForEvents("americanfootball_ncaaf", "us")
		Expect(err).NotTo(HaveOccurred())
		Expect(odds).NotTo(BeEmpty())
		Expect(len(odds)).To(BeNumerically(">=", 0))
		Expect(len(odds[0].Bookmakers)).To(BeNumerically(">=", 0))
		Expect(odds[0].Bookmakers[0].Key).NotTo(BeEmpty())
	})

})
