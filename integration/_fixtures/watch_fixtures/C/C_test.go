package C_test

import (
	. "$ROOT_PATH$/C"

	. "github.com/lsoica/ginkgo"
	. "github.com/lsoica/ginkgo"
)

var _ = Describe("C", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
