package books_test

import (
	"testing"

	. "github.com/lsoica/ginkgo"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}
