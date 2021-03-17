package codelocation_test

import (
	. "github.com/lsoica/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCodelocation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CodeLocation Suite")
}
