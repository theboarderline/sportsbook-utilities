package odds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/sportsbook-utilities/odds"
)

var _ = Describe("Bets", func() {

	It("can format the price string for an underdog bet", func() {
		bet := odds.Outcome{
			Name:  "Chicago Bulls",
			Price: 180,
		}

		betString := bet.FormatPrice()
		Expect(betString).To(Equal("Chicago Bulls +180"))
	})

	It("can format the price string for a favorite bet", func() {
		bet := odds.Outcome{
			Name:  "Chicago Bulls",
			Price: -180,
		}

		betString := bet.FormatPrice()
		Expect(betString).To(Equal("Chicago Bulls -180"))
	})

})
