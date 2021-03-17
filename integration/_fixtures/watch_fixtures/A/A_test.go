package A_test

import (
	. "$ROOT_PATH$/A"

	. "github.com/lsoica/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("A", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
