package C_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "C Suite")
}
