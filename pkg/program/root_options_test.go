package program

import (
	"github.com/franela/goblin"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandleInputDir(t *testing.T) {

	cwd, _ := os.Getwd()

	g := goblin.Goblin(t)

	g.Describe("Multiple cases for handleInputDir", func() {

		g.It("Should be cwd if argument is an empty string", func() {
			s, err := handleInputDir("")
			g.Assert(s).Equal(cwd)
			g.Assert(err).Equal(nil)
		})

		g.It("Should be abs if argument is relative", func() {
			s, err := handleInputDir("./example/of/relative/path")
			g.Assert(filepath.IsAbs(s)).Equal(true)
			g.Assert(strings.HasPrefix(s, cwd)).Equal(true)
			g.Assert(err).Equal(nil)
		})

		g.It("Should not change if argument is already abs", func() {
			s, err := handleInputDir("/example/of/absolute/path")
			g.Assert(filepath.IsAbs(s)).Equal(true)
			g.Assert(s).Equal("/example/of/absolute/path")
			g.Assert(err).Equal(nil)
		})

	})

}

func TestHandleOutputFile(t *testing.T) {

	cwd, _ := os.Getwd()

	g := goblin.Goblin(t)

	g.Describe("Multiple cases for handleOutputFile", func() {

		g.It("Should be __default_init_filename if argument is an empty string", func() {
			expected, _ := filepath.Abs(__default_init_filename)
			result, err := handleOutputFile("")
			g.Assert(result).Equal(expected)
			g.Assert(err).Equal(nil)
		})

		g.It("Should be abs if argument is relative", func() {
			s, err := handleOutputFile("./example/of/relative/path")
			g.Assert(filepath.IsAbs(s)).Equal(true)
			g.Assert(strings.HasPrefix(s, cwd)).Equal(true)
			g.Assert(err).Equal(nil)
		})

		g.It("Should not change if argument is already abs", func() {
			s, err := handleOutputFile("/example/of/absolute/path")
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
