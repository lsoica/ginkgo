package fail_fixture_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestFail_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fail_fixture Suite")
}
