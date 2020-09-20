package gen

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

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
			//			close(errChan)
			//			close(doneChan)
		}()

		iw.doWrite(iw.OutputFilename, filesChan, errChan)
	}()

	return errChan, doneChan
}

func (iw *InitWriter) doWrite(filename string, filesChan <-chan *File, errChan chan<- error) {

	// Create the init file.
	initFile, err := os.Create(filename)
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
	initFile.WriteString("\n")
	initFile.WriteString("func init() {\n")

	// Write every loaded file.
	for file := range filesChan {
		lineBeginning := fmt.Sprintf(`Files.db["%s"] = []byte{`, file.key)
		initFile.WriteString(lineBeginning)

		for _, b := range file.value {
			singleByte := fmt.Sprintf("%d,", b)
			initFile.WriteString(singleByte)
		}

		lineEnd := "}\n\n"
		initFile.WriteString(lineEnd)
	}

	// Close the init() function.
	initFile.WriteString("}\n")

	//
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
