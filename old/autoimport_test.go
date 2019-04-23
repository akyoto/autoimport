package autoimport_test

import (
	"testing"

	"github.com/akyoto/autoimport"
)

func TestCompile(t *testing.T) {
	files := []string{
		"examples/hello.go.txt",
	}

	autoimport.Files(files)
}
