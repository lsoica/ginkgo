package more_ginkgo_tests_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestMore_ginkgo_tests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "More_ginkgo_tests Suite")
}
