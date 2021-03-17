package types_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
	. "github.com/onsi/gomega"
)

func TestTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Types Suite")
}
