package A_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Suite")
}
