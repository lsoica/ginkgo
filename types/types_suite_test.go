package types_test

import (
	. "github.com/lsoica/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Types Suite")
}
