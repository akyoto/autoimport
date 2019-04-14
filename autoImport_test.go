package main

import "testing"

func TestCompile(t *testing.T) {
	files := []string{
		"examples/hello.go.txt",
	}

	autoImport(files)
}
