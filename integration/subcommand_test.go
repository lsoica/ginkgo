package integration_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "github.com/lsoica/ginkgo"
	"github.com/onsi/gomega/gexec"
	"github.com/lsoica/ginkgo/types"
	. "github.com/onsi/gomega"
)

var _ = Describe("Subcommand", func() {
	Describe("ginkgo bootstrap", func() {
		var pkgPath string
		BeforeEach(func() {
			pkgPath = tmpPath("foo")
			os.Mkdir(pkgPath, 0777)
		})

		It("should generate a bootstrap file, as long as one does not exist", func() {
			session := startGinkgo(pkgPath, "bootstrap")
			Eventually(session).Should(gexec.Exit(0))
			output := session.Out.Contents()

			Ω(output).Should(ContainSubstring("foo_suite_test.go"))

			content, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_suite_test.go"))
			Ω(err).ShouldNot(HaveOccurred())
			Ω(content).Should(ContainSubstring("package foo_test"))
			Ω(content).Should(ContainSubstring("func TestFoo(t *testing.T) {"))
			Ω(content).Should(ContainSubstring("RegisterFailHandler"))
			Ω(content).Should(ContainSubstring("RunSpecs"))

			Ω(content).Should(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))
			Ω(content).Should(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))

			session = startGinkgo(pkgPath, "bootstrap")
			Eventually(session).Should(gexec.Exit(1))
			output = session.Out.Contents()
			Ω(output).Should(ContainSubstring("foo_suite_test.go already exists"))
		})

		It("should import nodot declarations when told to", func() {
			session := startGinkgo(pkgPath, "bootstrap", "--nodot")
			Eventually(session).Should(gexec.Exit(0))
			output := session.Out.Contents()

			Ω(output).Should(ContainSubstring("foo_suite_test.go"))

			content, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_suite_test.go"))
			Ω(err).ShouldNot(HaveOccurred())
			Ω(content).Should(ContainSubstring("package foo_test"))
			Ω(content).Should(ContainSubstring("func TestFoo(t *testing.T) {"))
			Ω(content).Should(ContainSubstring("RegisterFailHandler"))
			Ω(content).Should(ContainSubstring("RunSpecs"))

			Ω(content).Should(ContainSubstring("var It = ginkgo.It"))
			Ω(content).Should(ContainSubstring("var Ω = gomega.Ω"))

			Ω(content).Should(ContainSubstring("\t" + `"github.com/lsoica/ginkgo"`))
			Ω(content).Should(ContainSubstring("\t" + `"github.com/lsoica/ginkgo"`))
		})

		It("should generate an agouti bootstrap file when told to", func() {
			session := startGinkgo(pkgPath, "bootstrap", "--agouti")
			Eventually(session).Should(gexec.Exit(0))
			output := session.Out.Contents()

			Ω(output).Should(ContainSubstring("foo_suite_test.go"))

			content, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_suite_test.go"))
			Ω(err).ShouldNot(HaveOccurred())
			Ω(content).Should(ContainSubstring("package foo_test"))
			Ω(content).Should(ContainSubstring("func TestFoo(t *testing.T) {"))
			Ω(content).Should(ContainSubstring("RegisterFailHandler"))
			Ω(content).Should(ContainSubstring("RunSpecs"))

			Ω(content).Should(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))
			Ω(content).Should(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))
			Ω(content).Should(ContainSubstring("\t" + `"github.com/sclevine/agouti"`))
		})

		It("should generate a bootstrap file using a template when told to", func() {
			templateFile := filepath.Join(pkgPath, ".bootstrap")
			ioutil.WriteFile(templateFile, []byte(`package {{.Package}}

			import (
				{{.GinkgoImport}}
				{{.GomegaImport}}

				"testing"
				"binary"
			)

			func Test{{.FormattedName}}(t *testing.T) {
				// This is a {{.Package}} test
			}`), 0666)
			session := startGinkgo(pkgPath, "bootstrap", "--template", ".bootstrap")
			Eventually(session).Should(gexec.Exit(0))
			output := session.Out.Contents()

			Ω(output).Should(ContainSubstring("foo_suite_test.go"))

			content, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_suite_test.go"))
			Ω(err).ShouldNot(HaveOccurred())
			Ω(content).Should(ContainSubstring("package foo_test"))
			Ω(content).Should(ContainSubstring(`. "github.com/lsoica/ginkgo"`))
			Ω(content).Should(ContainSubstring(`. "github.com/lsoica/ginkgo"`))
			Ω(content).Should(ContainSubstring(`"binary"`))
			Ω(content).Should(ContainSubstring("// This is a foo_test test"))
		})
	})

	Describe("nodot", func() {
		It("should update the declarations in the bootstrap file", func() {
			pkgPath := tmpPath("foo")
			os.Mkdir(pkgPath, 0777)

			session := startGinkgo(pkgPath, "bootstrap", "--nodot")
			Eventually(session).Should(gexec.Exit(0))

			byteContent, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_suite_test.go"))
			Ω(err).ShouldNot(HaveOccurred())

			content := string(byteContent)
			content = strings.Replace(content, "var It =", "var MyIt =", -1)
			content = strings.Replace(content, "var Ω = gomega.Ω\n", "", -1)

			err = ioutil.WriteFile(filepath.Join(pkgPath, "foo_suite_test.go"), []byte(content), os.ModePerm)
			Ω(err).ShouldNot(HaveOccurred())

			session = startGinkgo(pkgPath, "nodot")
			Eventually(session).Should(gexec.Exit(0))

			byteContent, err = ioutil.ReadFile(filepath.Join(pkgPath, "foo_suite_test.go"))
			Ω(err).ShouldNot(HaveOccurred())

			Ω(byteContent).Should(ContainSubstring("var MyIt = ginkgo.It"))
			Ω(byteContent).ShouldNot(ContainSubstring("var It = ginkgo.It"))
			Ω(byteContent).Should(ContainSubstring("var Ω = gomega.Ω"))
		})
	})

	Describe("ginkgo generate", func() {
		var pkgPath string

		BeforeEach(func() {
			pkgPath = tmpPath("foo_bar")
			os.Mkdir(pkgPath, 0777)
		})

		Context("with no arguments", func() {
			It("should generate a test file named after the package", func() {
				session := startGinkgo(pkgPath, "generate")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("foo_bar_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_bar_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`var _ = Describe("FooBar", func() {`))
				Ω(content).Should(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))
				Ω(content).Should(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))

				session = startGinkgo(pkgPath, "generate")
				Eventually(session).Should(gexec.Exit(1))
				output = session.Out.Contents()

				Ω(output).Should(ContainSubstring("foo_bar_test.go already exists"))
			})
		})

		Context("with template argument", func() {
			It("should generate a test file using a template", func() {
				templateFile := filepath.Join(pkgPath, ".generate")
				ioutil.WriteFile(templateFile, []byte(`package {{.Package}}
				import (
					{{if .IncludeImports}}. "github.com/lsoica/ginkgo"{{end}}
					{{if .IncludeImports}}. "github.com/lsoica/ginkgo"{{end}}
				
					{{if .ImportPackage}}"{{.PackageImportPath}}"{{end}}
				)
				
				var _ = Describe("{{.Subject}}", func() {
					// This is a {{.Package}} test
				})`), 0666)
				session := startGinkgo(pkgPath, "generate", "--template", ".generate")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("foo_bar_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_bar_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`. "github.com/lsoica/ginkgo"`))
				Ω(content).Should(ContainSubstring(`. "github.com/lsoica/ginkgo"`))
				Ω(content).Should(ContainSubstring(`/foo_bar"`))
				Ω(content).Should(ContainSubstring("// This is a foo_bar_test test"))
			})
		})

		Context("with an argument of the form: foo", func() {
			It("should generate a test file named after the argument", func() {
				session := startGinkgo(pkgPath, "generate", "baz_buzz")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("baz_buzz_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "baz_buzz_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`var _ = Describe("BazBuzz", func() {`))
			})
		})

		Context("with an argument of the form: foo.go", func() {
			It("should generate a test file named after the argument", func() {
				session := startGinkgo(pkgPath, "generate", "baz_buzz.go")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("baz_buzz_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "baz_buzz_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`var _ = Describe("BazBuzz", func() {`))

			})
		})

		Context("with an argument of the form: foo_test", func() {
			It("should generate a test file named after the argument", func() {
				session := startGinkgo(pkgPath, "generate", "baz_buzz_test")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("baz_buzz_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "baz_buzz_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`var _ = Describe("BazBuzz", func() {`))
			})
		})

		Context("with an argument of the form: foo-test", func() {
			It("should generate a test file named after the argument", func() {
				session := startGinkgo(pkgPath, "generate", "baz-buzz-test")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("baz_buzz_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "baz_buzz_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`var _ = Describe("BazBuzz", func() {`))
			})
		})

		Context("with an argument of the form: foo_test.go", func() {
			It("should generate a test file named after the argument", func() {
				session := startGinkgo(pkgPath, "generate", "baz_buzz_test.go")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("baz_buzz_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "baz_buzz_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`var _ = Describe("BazBuzz", func() {`))
			})
		})

		Context("with multiple arguments", func() {
			It("should generate a test file named after the argument", func() {
				session := startGinkgo(pkgPath, "generate", "baz", "buzz")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("baz_test.go"))
				Ω(output).Should(ContainSubstring("buzz_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "baz_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`var _ = Describe("Baz", func() {`))

				content, err = ioutil.ReadFile(filepath.Join(pkgPath, "buzz_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring(`var _ = Describe("Buzz", func() {`))
			})
		})

		Context("with nodot", func() {
			It("should not import ginkgo or gomega", func() {
				session := startGinkgo(pkgPath, "generate", "--nodot")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("foo_bar_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_bar_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).ShouldNot(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))
				Ω(content).ShouldNot(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))
			})
		})

		Context("with agouti", func() {
			It("should generate an agouti test file", func() {
				session := startGinkgo(pkgPath, "generate", "--agouti")
				Eventually(session).Should(gexec.Exit(0))
				output := session.Out.Contents()

				Ω(output).Should(ContainSubstring("foo_bar_test.go"))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "foo_bar_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package foo_bar_test"))
				Ω(content).Should(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))
				Ω(content).Should(ContainSubstring("\t" + `. "github.com/lsoica/ginkgo"`))
				Ω(content).Should(ContainSubstring("\t" + `. "github.com/sclevine/agouti/matchers"`))
				Ω(content).Should(ContainSubstring("\t" + `"github.com/sclevine/agouti"`))
				Ω(content).Should(ContainSubstring("page, err = agoutiDriver.NewPage()"))
			})
		})
	})

	Describe("ginkgo bootstrap/generate", func() {
		var pkgPath string
		BeforeEach(func() {
			pkgPath = tmpPath("some-crazy-thing")
			os.Mkdir(pkgPath, 0777)
		})

		Context("when the working directory is empty", func() {
			It("generates correctly named bootstrap and generate files with a package name derived from the directory", func() {
				session := startGinkgo(pkgPath, "bootstrap")
				Eventually(session).Should(gexec.Exit(0))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "some_crazy_thing_suite_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package some_crazy_thing_test"))
				Ω(content).Should(ContainSubstring("SomeCrazyThing Suite"))

				session = startGinkgo(pkgPath, "generate")
				Eventually(session).Should(gexec.Exit(0))

				content, err = ioutil.ReadFile(filepath.Join(pkgPath, "some_crazy_thing_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package some_crazy_thing_test"))
				Ω(content).Should(ContainSubstring("SomeCrazyThing"))
			})
		})

		Context("when the working directory contains a file with a package name", func() {
			BeforeEach(func() {
				Ω(ioutil.WriteFile(filepath.Join(pkgPath, "foo.go"), []byte("package main\n\nfunc main() {}"), 0777)).Should(Succeed())
			})

			It("generates correctly named bootstrap and generate files with the package name", func() {
				session := startGinkgo(pkgPath, "bootstrap")
				Eventually(session).Should(gexec.Exit(0))

				content, err := ioutil.ReadFile(filepath.Join(pkgPath, "some_crazy_thing_suite_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package main_test"))
				Ω(content).Should(ContainSubstring("SomeCrazyThing Suite"))

				session = startGinkgo(pkgPath, "generate")
				Eventually(session).Should(gexec.Exit(0))

				content, err = ioutil.ReadFile(filepath.Join(pkgPath, "some_crazy_thing_test.go"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(content).Should(ContainSubstring("package main_test"))
				Ω(content).Should(ContainSubstring("SomeCrazyThing"))
			})
		})
	})

	Describe("Go module and sginkgo bootstrap/generate", func() {
		var (
			pkgPath     string
			savedGoPath string
		)

		BeforeEach(func() {
			pkgPath = tmpPath("myamazingmodule")
			os.Mkdir(pkgPath, 0777)
			Expect(ioutil.WriteFile(filepath.Join(pkgPath, "go.mod"), []byte("module fake.com/me/myamazingmodule\n"), 0777)).To(Succeed())
			savedGoPath = os.Getenv("GOPATH")
			Expect(os.Setenv("GOPATH", "")).To(Succeed())
			Expect(os.Setenv("GO111MODULE", "on")).To(Succeed()) // needed pre-Go 1.13
		})

		AfterEach(func() {
			Expect(os.Setenv("GOPATH", savedGoPath)).To(Succeed())
			Expect(os.Setenv("GO111MODULE", "")).To(Succeed())
		})

		It("generates correctly named bootstrap and generate files with the module name", func() {
			session := startGinkgo(pkgPath, "bootstrap")
			Eventually(session).Should(gexec.Exit(0))

			content, err := ioutil.ReadFile(filepath.Join(pkgPath, "myamazingmodule_suite_test.go"))
			Expect(err).NotTo(HaveOccurred())
			Expect(content).To(ContainSubstring("package myamazingmodule_test"), string(content))
			Expect(content).To(ContainSubstring("Myamazingmodule Suite"), string(content))

			session = startGinkgo(pkgPath, "generate")
			Eventually(session).Should(gexec.Exit(0))

			content, err = ioutil.ReadFile(filepath.Join(pkgPath, "myamazingmodule_test.go"))
			Expect(err).NotTo(HaveOccurred())
			Expect(content).To(ContainSubstring("package myamazingmodule_test"), string(content))
			Expect(content).To(ContainSubstring("fake.com/me/myamazingmodule"), string(content))
			Expect(content).To(ContainSubstring("Myamazingmodule"), string(content))
		})
	})

	Describe("ginkgo blur", func() {
		It("should unfocus tests", func() {
			pathToTest := tmpPath("focused")
			fixture := fixturePath("focused_fixture")
			copyIn(fixture, pathToTest, true)

			session := startGinkgo(pathToTest, "--noColor", "-r")
			Eventually(session).Should(gexec.Exit(types.GINKGO_FOCUS_EXIT_CODE))
			output := session.Out.Contents()

			Ω(string(output)).Should(ContainSubstring("Detected Programmatic Focus"))

			session = startGinkgo(pathToTest, "blur")
			Eventually(session).Should(gexec.Exit(0))
			output = session.Out.Contents()
			Ω(string(output)).ShouldNot(ContainSubstring("expected 'package'"))

			session = startGinkgo(pathToTest, "--noColor", "-r")
			Eventually(session).Should(gexec.Exit(0))
			output = session.Out.Contents()
			Ω(string(output)).Should(ContainSubstring("Ginkgo ran 2 suites"))
			Ω(string(output)).Should(ContainSubstring("Test Suite Passed"))
			Ω(string(output)).ShouldNot(ContainSubstring("Detected Programmatic Focus"))

			Expect(sameFile(filepath.Join(pathToTest, "README.md"), filepath.Join(fixture, "README.md"))).To(BeTrue())
		})

		It("should ignore the 'vendor' folder", func() {
			pathToTest := tmpPath("focused_fixture_with_vendor")
			copyIn(fixturePath("focused_fixture_with_vendor"), pathToTest, true)

			session := startGinkgo(pathToTest, "blur")
			Eventually(session).Should(gexec.Exit(0))

			session = startGinkgo(pathToTest, "--noColor")
			Eventually(session).Should(gexec.Exit(0))
			output := session.Out.Contents()
			Expect(string(output)).To(ContainSubstring("13 Passed"))
			Expect(string(output)).To(ContainSubstring("0 Skipped"))

			vendorPath := fixturePath("focused_fixture_with_vendor/vendor")
			otherVendorPath := filepath.Join(pathToTest, "vendor")

			Expect(sameFolder(vendorPath, otherVendorPath)).To(BeTrue())
		})
	})

	Describe("ginkgo version", func() {
		It("should print out the version info", func() {
			session := startGinkgo("", "version")
			Eventually(session).Should(gexec.Exit(0))
			output := session.Out.Contents()

			Ω(output).Should(MatchRegexp(`Ginkgo Version \d+\.\d+\.\d+`))
		})
	})

	Describe("ginkgo help", func() {
		It("should print out usage information", func() {
			session := startGinkgo("", "help")
			Eventually(session).Should(gexec.Exit(0))
			output := string(session.Out.Contents())

			Ω(output).Should(MatchRegexp(`Ginkgo Version \d+\.\d+\.\d+`))
			Ω(output).Should(ContainSubstring("ginkgo watch"))
			Ω(output).Should(ContainSubstring("-succinct"))
			Ω(output).Should(ContainSubstring("-nodes"))
			Ω(output).Should(ContainSubstring("ginkgo generate"))
			Ω(output).Should(ContainSubstring("ginkgo help <COMMAND>"))
		})
	})
})
