package second_package_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestCoverageFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CombinedFixture Second Suite")
}
