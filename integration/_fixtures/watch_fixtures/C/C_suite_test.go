package C_test

import (
	. "github.com/lsoica/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "C Suite")
}
