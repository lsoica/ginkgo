package failing_ginkgo_tests_test

import (
	. "github.com/lsoica/ginkgo"
	. "github.com/lsoica/ginkgo/integration/_fixtures/failing_ginkgo_tests"
)

var _ = Describe("FailingGinkgoTests", func() {
	It("should fail", func() {
		Ω(AlwaysFalse()).Should(BeTrue())
	})

	It("should pass", func() {
		Ω(AlwaysFalse()).Should(BeFalse())
	})
})
