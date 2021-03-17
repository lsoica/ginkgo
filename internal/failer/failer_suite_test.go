package failer_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
	. "github.com/onsi/gomega"
)

func TestFailer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Failer Suite")
}
