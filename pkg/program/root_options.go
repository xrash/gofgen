package program

import (
	"os"
	"path/filepath"
)

const __default_init_filename = "./init_gofgen.go"

func handleInputDir(s string) (string, error) {
	if s == "" {
		return os.Getwd()
	}

	return filepath.Abs(s)
}

func handleOutputFile(s string) (string, error) {
	if s == "" {
		s = __default_init_filename
	}

	return filepath.Abs(s)
}

func handlePackageName(packageName, inputDirname string) string {
	if packageName == "" {
		return filepath.Base(inputDirname)
	}

	return packageName
}

func handleCompress(s string) (string, error) {
	return "", nil
}
