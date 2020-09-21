[![Build Status](https://travis-ci.org/xrash/gofgen.svg?branch=master)](http://travis-ci.org/xrash/gofgen)

# gofgen

## What is this about

`gofgen` transforms a directory into a usable Go package that exposes the files under that directory to your program at runtime.

## Motivation

Suppose that in your Go project, you have a directory called `res` (it could have any name), containing static files:

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

Before compiling your program, you `cd` into `res` and run `gofgen`. `gofgen`, then, generates a file named `init_gofgen.go`. This file contains Go code loading everything under `res` into memory at compile time.

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

Finally, you can access all the files under `res` from your code. For example, `main.go` might look like this:

```go
package main

import (
	"fmt"
	"<yourpackage>/res"
)

func main() {
	b, ok := res.FS.Get("/html/index.html")
	fmt.Println(string(b))
	fmt.Println(ok)
}
```

As the example shows, you can access the in-memory file system, loaded by gofgen's init file, using the `FS` variable now available from your `res` package. Every file inside `res` must be referred from the root directory, that is, starting with an `/`, as in `/html/index.html`.

Your directory `res` was once just a local storage for your static files - now it's an accessible Go package!

## Examples

### using all options

```bash
$ gofgen --package-name main --input-dir /movies/taxi-driver/screenshots --output-file /opt/some-amazing-go-project/init_screenshots.go
```

This will generate `/opt/some-amazing-go-project/init.go`, having `package main`, containing Go code that exposes the files under `/movies/taxi-driver/screenshots`.

### using zero options

```bash
$ cd ~/my-amazing-project/resources && gofgen
```

This will generate `~/my-amazing-project/resources/init_gofgen.go`, having `package resources`, exposing files under `~/my-amazing-project/resources`.

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

```
gofgen generates go code loading your local files into memory at compile time

Usage:
  gofgen [flags]

Flags:
  -h, --help                  help for gofgen
      --input-dir string      dirname of the directory to read from (default ".")
      --output-file string    filename of the output file (default "./init_gofgen.go")
      --package-name string   package name to use, default value is basename of input-dirname
```
