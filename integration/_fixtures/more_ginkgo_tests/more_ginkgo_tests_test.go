package more_ginkgo_tests_test

import (
	. "github.com/lsoica/ginkgo"
	. "github.com/lsoica/ginkgo/integration/_fixtures/more_ginkgo_tests"
)

var _ = Describe("MoreGinkgoTests", func() {
	It("should pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})

	It("should always pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})
})
