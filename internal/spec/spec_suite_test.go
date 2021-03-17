package spec_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Suite")
}
