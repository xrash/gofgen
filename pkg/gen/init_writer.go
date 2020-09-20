package gen

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

func osCreateFile(filename string) (io.WriteCloser, error) {
	return os.Create(filename)
}

var createFile func(string) (io.WriteCloser, error) = osCreateFile

type InitWriter struct {
	PackageName          string
	CompressionAlgorithm string
	OutputFilename       string
}

func (iw *InitWriter) WriteFuture(filesChan <-chan *File) (<-chan error, <-chan bool) {
	errChan := make(chan error)
	doneChan := make(chan bool)

	go func() {
		defer func() {
			doneChan <- true
		}()

		iw.doWrite(iw.OutputFilename, filesChan, errChan)
	}()

	return errChan, doneChan
}

func (iw *InitWriter) doWrite(filename string, filesChan <-chan *File, errChan chan<- error) {

	// Create the init file.
	initFile, err := createFile(filename)
	if err != nil {
		errChan <- fmt.Errorf("Error creating init file %s: %v", filename, err)
		return
	}

	defer func() {
		if err := initFile.Close(); err != nil {
			errChan <- err
			return
		}
	}()

	// Write the head of the init file.
	if err := iw.writeHead(initFile); err != nil {
		errChan <- fmt.Errorf("Error writing head to init file %s: %v", filename, err)
		return
	}

	// Write the beginning of the init() function.
	initFile.Write([]byte("\n"))
	initFile.Write([]byte("func init() {\n\n"))

	// Write every loaded file.
	for file := range filesChan {
		lineBeginning := fmt.Sprintf(`Files.db["%s"] = []byte{`, file.key)
		initFile.Write([]byte(lineBeginning))

		for _, b := range file.value {
			singleByte := fmt.Sprintf("%d,", b)
			initFile.Write([]byte(singleByte))
		}

		lineEnd := "}\n\n"
		initFile.Write([]byte(lineEnd))
	}

	// Close the init() function.
	initFile.Write([]byte("}\n"))
}

func (iw *InitWriter) writeHead(w io.Writer) error {
	tmpl, err := template.New("init_file_template_head").Parse(__init_file_template_head)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"packageName": iw.PackageName,
	}

	return tmpl.Execute(w, data)
}
