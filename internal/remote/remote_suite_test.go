package remote_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
	. "github.com/onsi/gomega"
)

func TestRemote(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Remote Spec Forwarding Suite")
}
