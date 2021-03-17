package testingtproxy_test

import (
	"testing"

	. "github.com/lsoica/ginkgo"
)

func TestTestingtproxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testingtproxy Suite")
}
