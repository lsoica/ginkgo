package D_test

import (
	. "$ROOT_PATH$/D"

	. "github.com/lsoica/ginkgo"
	. "github.com/lsoica/ginkgo"
)

var _ = Describe("D", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
