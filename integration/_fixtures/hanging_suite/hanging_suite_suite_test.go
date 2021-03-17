package hanging_suite_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestHangingSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HangingSuite Suite")
}
