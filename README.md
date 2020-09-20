[![Build Status](https://travis-ci.org/xrash/gofgen.svg?branch=master)](http://travis-ci.org/xrash/gofgen)

# gofgen

## What is this about

In your Go project, you have a directory `res` containing static files:

```
project/
├── main.go
└── res/
    ├── html/
    |   └── index.html
    └── css/
        ├── main.css
        └── utils.css
```

You run `gofgen` inside `project/res` before compiling your program. `gofgen` then generates file `project/res/init_gofgen.go`, containing Go code that loads everything under `project/res` into memory.

```
project/
├── main.go
└── res/
    ├── html/
    |   └── index.html
    ├── css/
    |   ├── main.css
    |   └── utils.css
    └── init_gofgen.go
```

Finally, you can access all the files under `project/res` from your code. For example, `main.go` might look like this:

```go
package main

import (
	"fmt"
	"github.com/xrash/testgofgen/res"
)

func main() {
	b, ok := res.FS.Get("/html/index.html")
	fmt.Println(string(b))
	fmt.Println(ok)
}
```

As the example shows, you can access the in-memory file system, loaded by gofgen's init file, using the `FS` variable now available from your `res` package. Every file inside `res` must be referred from the root directory, that is, starting with an `/`, as in `/html/index.html`.

Your directory `res` was once just a local storage for your static files - now it's an accessible Go package!

## How to install

Using `go get`:

```bash
go get github.com/xrash/gofgen/cmd/gofgen
```

Manually compiling:

```bash
$ git clone https://github.com/xrash/gofgen.git
$ cd gofgen
$ make
$ sudo mv bin/gofgen /usr/local/bin
```

## Usage

```bash
gofgen generates go code loading your local files into memory at compile time

Usage:
  gofgen [flags]

Flags:
  -h, --help                     help for gofgen
      --input-dirname string     dirname of the directory to read from (default ".")
      --output-filename string   filename of the output file (default "./init_gofgen.go")
      --package-name string      package name to use, default value is basename of input-dirname
```
