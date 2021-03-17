package flags_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestFlags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flags Suite")
}
