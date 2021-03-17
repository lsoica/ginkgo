package nested_test

import (
	"testing"

	. "github.com/lsoica/ginkgo"
)

func TestNested(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nested Suite")
}
