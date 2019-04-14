package autoimport_test

import "testing"

func TestCompile(t *testing.T) {
	files := []string{
		"examples/hello.go.txt",
	}

	autoimport.Files(files)
}
