package spec_iterator_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestSpecIterator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SpecIterator Suite")
}
