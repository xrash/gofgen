package gen

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type FilesReader struct {
	FilenamesToIgnore []string
	RootDirname       string
}

func (fr *FilesReader) ReadFuture(filesChan chan<- *File) <-chan error {
	errChan := make(chan error)

	go func() {
		defer func() {
			//			close(errChan)
			close(filesChan)
		}()

		fr.readDir(fr.RootDirname, filesChan, errChan)
	}()

	return errChan
}

func (fr *FilesReader) readDir(dirname string, filesChan chan<- *File, errChan chan<- error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		errChan <- fmt.Errorf("Error reading dir %s: %v", dirname, err)
		return
	}

	for _, finfo := range files {
		filename := filepath.Join(dirname, finfo.Name())
		relativeFilename := filename[len(fr.RootDirname):]

		if fr.shouldIgnore(relativeFilename) {
			continue
		}

		if finfo.Mode().IsDir() {
			fr.readDir(filename, filesChan, errChan)
			continue
		}

		b, err := ioutil.ReadFile(filename)
		if err != nil {
			errChan <- fmt.Errorf("Error reading file %s: %v", filename, err)
			return
		}

		filesChan <- &File{
			key:   relativeFilename,
			value: b,
		}

	}
}

func (fr *FilesReader) shouldIgnore(filename string) bool {
	for _, filenameToIgnore := range fr.FilenamesToIgnore {
		if filenameToIgnore == filename {
			return true
		}
	}

	return false
}
