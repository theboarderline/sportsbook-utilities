package odds_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSportsbook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sportsbook Suite")
}
