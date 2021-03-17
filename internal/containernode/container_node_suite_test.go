package containernode_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestContainernode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containernode Suite")
}
