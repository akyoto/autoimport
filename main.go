package main

import (
	"flag"
)

var packagesByName = map[string][]*Package{}

func main() {
	flag.Parse()
	files := flag.Args()
	autoImport(files)
}
