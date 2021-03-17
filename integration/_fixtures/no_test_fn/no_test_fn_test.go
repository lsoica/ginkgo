package no_test_fn_test

import (
	. "github.com/lsoica/ginkgo"
	. "github.com/lsoica/ginkgo/integration/_fixtures/no_test_fn"
)

var _ = Describe("NoTestFn", func() {
	It("should proxy strings", func() {
		Ω(StringIdentity("foo")).Should(Equal("foo"))
	})
})
