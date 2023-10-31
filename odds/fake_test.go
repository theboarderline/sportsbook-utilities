package odds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/sportsbook-utilities/odds"
)

var _ = Describe("Fake", func() {

	var (
		sportsbookClient *odds.FakeClient
	)

	BeforeEach(func() {
		sportsbookClient = odds.NewFakeClient()
	})

	It("can get the upcoming sporting leagues from the fake client", func() {
		leagues, err := sportsbookClient.GetSportingLeagues()
		Expect(err).NotTo(HaveOccurred())
		Expect(leagues).NotTo(BeEmpty())
		Expect(len(leagues)).To(BeNumerically(">=", 0))
	})

	It("can get the upcoming sporting events from the fake client", func() {
		events, err := sportsbookClient.GetUpcomingEvents("americanfootball_ncaaf")
		Expect(err).NotTo(HaveOccurred())
		Expect(events).NotTo(BeEmpty())
		Expect(len(events)).To(BeNumerically(">=", 0))
	})

	It("can get the odds for a game from the fake client", func() {
		odds, err := sportsbookClient.GetOddsForEvents("americanfootball_ncaaf", []string{"us"}, []string{"h2h"})
		Expect(err).NotTo(HaveOccurred())
		Expect(odds).NotTo(BeEmpty())
		Expect(len(odds)).To(BeNumerically(">=", 0))
		Expect(len(odds[0].Bookmakers)).To(BeNumerically(">=", 0))
		Expect(odds[0].Bookmakers[0].Key).NotTo(BeEmpty())
		Expect(odds[0].Bookmakers[0].Markets).NotTo(BeEmpty())
		Expect(odds[0].Bookmakers[0].Markets[0].Outcomes).NotTo(BeEmpty())
		Expect(len(odds[0].Bookmakers[0].Markets[0].Outcomes)).To(BeNumerically(">=", 0))
		Expect(odds[0].Bookmakers[0].Markets[0].Outcomes[0].Name).NotTo(BeEmpty())
		Expect(odds[0].Bookmakers[0].Markets[0].Outcomes[0].Price).NotTo(BeZero())
	})

})
