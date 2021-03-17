package tags_tests_test

import (
	. "github.com/lsoica/ginkgo"

	"testing"
)

func TestTagsTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TagsTests Suite")
}
