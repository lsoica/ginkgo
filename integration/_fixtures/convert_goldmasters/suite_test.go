package tmp_test

import (
	"testing"

	. "github.com/lsoica/ginkgo"
)

func TestConvertFixtures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConvertFixtures Suite")
}
