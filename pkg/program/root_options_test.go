package program

import (
	"github.com/franela/goblin"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandleInputDirname(t *testing.T) {

	cwd, _ := os.Getwd()

	g := goblin.Goblin(t)

	g.Describe("Multiple cases for handleInputDirname", func() {

		g.It("Should be cwd if argument is an empty string", func() {
			s, err := handleInputDirname("")
			g.Assert(s).Equal(cwd)
			g.Assert(err).Equal(nil)
		})

		g.It("Should be abs if argument is relative", func() {
			s, err := handleInputDirname("./example/of/relative/path")
			g.Assert(filepath.IsAbs(s)).Equal(true)
			g.Assert(strings.HasPrefix(s, cwd)).Equal(true)
			g.Assert(err).Equal(nil)
		})

		g.It("Should not change if argument is already abs", func() {
			s, err := handleInputDirname("/example/of/absolute/path")
			g.Assert(filepath.IsAbs(s)).Equal(true)
			g.Assert(s).Equal("/example/of/absolute/path")
			g.Assert(err).Equal(nil)
		})

	})

}

func TestHandleOutputFilename(t *testing.T) {

	cwd, _ := os.Getwd()

	g := goblin.Goblin(t)

	g.Describe("Multiple cases for handleOutputFilename", func() {

		g.It("Should be __default_init_filename if argument is an empty string", func() {
			expected, _ := filepath.Abs(__default_init_filename)
			result, err := handleOutputFilename("")
			g.Assert(result).Equal(expected)
			g.Assert(err).Equal(nil)
		})

		g.It("Should be abs if argument is relative", func() {
			s, err := handleOutputFilename("./example/of/relative/path")
			g.Assert(filepath.IsAbs(s)).Equal(true)
			g.Assert(strings.HasPrefix(s, cwd)).Equal(true)
			g.Assert(err).Equal(nil)
		})

		g.It("Should not change if argument is already abs", func() {
			s, err := handleOutputFilename("/example/of/absolute/path")
			g.Assert(filepath.IsAbs(s)).Equal(true)
			g.Assert(s).Equal("/example/of/absolute/path")
			g.Assert(err).Equal(nil)
		})

	})

}

func TestPackageName(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Multiple cases for handlePackageName", func() {

		g.It("Should be basename of inputDirname if packageName is empty", func() {
			s := handlePackageName("", "/imagine/random/path/here")
			g.Assert(s).Equal("here")
		})

		g.It("Should be packageName if packageName isn't empty", func() {
			s := handlePackageName("whateverpkg", "/imagine/random/path/here")
			g.Assert(s).Equal("whateverpkg")
		})

	})

}
