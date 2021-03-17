package progress_fixture_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestProgressFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProgressFixture Suite")
}
