package B_test

import (
	. "$ROOT_PATH$/B"

	. "github.com/lsoica/ginkgo"
	. "github.com/lsoica/ginkgo"
)

var _ = Describe("B", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
