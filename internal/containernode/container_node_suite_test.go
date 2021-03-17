package containernode_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
	. "github.com/onsi/gomega"
)

func TestContainernode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containernode Suite")
}
