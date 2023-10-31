package odds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/sportsbook-utilities/odds"
)

var _ = Describe("Bets", func() {

	It("can format the string for an underdog bet", func() {
		bet := odds.Outcome{
			Name:  "Chicago Bulls",
			Point: 180,
		}

		betString := bet.Format()
		Expect(betString).To(Equal("Chicago Bulls +180"))
	})

	It("can format the string for a favorite bet", func() {
		bet := odds.Outcome{
			Name:  "Chicago Bulls",
			Point: -180,
		}

		betString := bet.Format()
		Expect(betString).To(Equal("Chicago Bulls -180"))
	})

})
