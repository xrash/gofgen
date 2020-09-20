package program

import (
	"fmt"
	"os"
)

func PrintlnErr(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

func PrintfErr(f string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, f, args...)
}
