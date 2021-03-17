package B_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "B Suite")
}
