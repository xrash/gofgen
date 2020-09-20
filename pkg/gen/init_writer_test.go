package gen

import (
	"bytes"
	"github.com/franela/goblin"
	"io"
	"strings"
	"testing"
)

type dumbWriteCloser struct {
	buf *bytes.Buffer
}

func (dwc *dumbWriteCloser) Write(b []byte) (int, error) {
	return dwc.buf.Write(b)
}

func (dwc *dumbWriteCloser) Close() error {
	return nil
}

var dumbBuffer *dumbWriteCloser = &dumbWriteCloser{
	buf: bytes.NewBuffer(nil),
}

func dumbCreateFile(filename string) (io.WriteCloser, error) {
	return dumbBuffer, nil
}

func TestWriteFuture(t *testing.T) {

	g := goblin.Goblin(t)

	createFile = dumbCreateFile

	g.Describe("Multiple tests for InputWriter", func() {

		g.It("Should write the init file correctly", func() {
			iw := &InitWriter{
				PackageName:          "pkgname",
				CompressionAlgorithm: "",
				OutputFilename:       "/any/path",
			}

			filesChan := make(chan *File, 1024)
			writerErr, writerDone := iw.WriteFuture(filesChan)

			filesChan <- &File{
				key:   "/suppose/this/is/a/file",
				value: []byte{1, 2, 3, 4, 5, 10, 50, 100},
			}
			filesChan <- &File{
				key:   "/suppose/this/is/another/file.html",
				value: []byte{150, 150, 150},
			}
			filesChan <- &File{
				key:   "/example.html",
				value: []byte{},
			}

			close(filesChan)

			select {
			case <-writerErr:
				g.Assert(true).Equal(false)
			case <-writerDone:
				g.Assert(true).Equal(true)
			}

			expected :=
				`
package pkgname

type FileSystem struct {
db map[string][]byte
}

var FS *FileSystem = &FileSystem{
db: make(map[string][]byte),
}

func (fs *FileSystem) Get(key string) ([]byte, bool) {
v, ok := fs.db[key]
return v, ok
}


func init() {

FS.db["/suppose/this/is/a/file"] = []byte{1,2,3,4,5,10,50,100,}

FS.db["/suppose/this/is/another/file.html"] = []byte{150,150,150,}

FS.db["/example.html"] = []byte{}

}
`

			result := strings.TrimSpace(dumbBuffer.buf.String())
			expected = strings.TrimSpace(expected)

			g.Assert(result).Equal(expected)

		})

	})

}
