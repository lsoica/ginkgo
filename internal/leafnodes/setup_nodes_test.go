package leafnodes_test

import (
	. "github.com/lsoica/ginkgo"
	"github.com/lsoica/ginkgo/types"

	. "github.com/lsoica/ginkgo/internal/leafnodes"

	"github.com/lsoica/ginkgo/internal/codelocation"
	. "github.com/onsi/gomega"
)

var _ = Describe("Setup Nodes", func() {
	Describe("BeforeEachNodes", func() {
		It("should report the correct type and code location", func() {
			codeLocation := codelocation.New(0)
			beforeEach := NewBeforeEachNode(func() {}, codeLocation, 0, nil, 3)
			Ω(beforeEach.Type()).Should(Equal(types.SpecComponentTypeBeforeEach))
			Ω(beforeEach.CodeLocation()).Should(Equal(codeLocation))
		})
	})

	Describe("AfterEachNodes", func() {
		It("should report the correct type and code location", func() {
			codeLocation := codelocation.New(0)
			afterEach := NewAfterEachNode(func() {}, codeLocation, 0, nil, 3)
			Ω(afterEach.Type()).Should(Equal(types.SpecComponentTypeAfterEach))
			Ω(afterEach.CodeLocation()).Should(Equal(codeLocation))
		})
	})

	Describe("JustBeforeEachNodes", func() {
		It("should report the correct type and code location", func() {
			codeLocation := codelocation.New(0)
			justBeforeEach := NewJustBeforeEachNode(func() {}, codeLocation, 0, nil, 3)
			Ω(justBeforeEach.Type()).Should(Equal(types.SpecComponentTypeJustBeforeEach))
			Ω(justBeforeEach.CodeLocation()).Should(Equal(codeLocation))
		})
	})
	Describe("JustAfterEachNodes", func() {
		It("should report the correct type and code location", func() {
			codeLocation := codelocation.New(0)
			justAfterEach := NewJustAfterEachNode(func() {}, codeLocation, 0, nil, 3)
			Ω(justAfterEach.Type()).Should(Equal(types.SpecComponentTypeJustAfterEach))
			Ω(justAfterEach.CodeLocation()).Should(Equal(codeLocation))
		})
	})
})
