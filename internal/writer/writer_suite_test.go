package writer_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestWriter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Writer Suite")
}
