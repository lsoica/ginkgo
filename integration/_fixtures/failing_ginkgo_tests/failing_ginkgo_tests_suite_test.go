package failing_ginkgo_tests_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestFailing_ginkgo_tests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Failing_ginkgo_tests Suite")
}
