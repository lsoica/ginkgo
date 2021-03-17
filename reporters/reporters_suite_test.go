package reporters_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestReporters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reporters Suite")
}
