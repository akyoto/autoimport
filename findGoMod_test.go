package autoimport

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/akyoto/assert"
)

func TestFindGoMod(t *testing.T) {
	currentDir, _ := filepath.Abs(".")
	files := map[string]string{
		"examples/hello.go.txt":         path.Join(currentDir, "go.mod"),
		"examples/nested/nested.go.txt": path.Join(currentDir, "go.mod"),
		"./":                            path.Join(currentDir, "go.mod"),
		".":                             path.Join(currentDir, "go.mod"),
		"":                              path.Join(currentDir, "go.mod"),
		"../":                           "",
		"..":                            "",
	}

	for file, result := range files {
		goModPath := findGoMod(file)
		assert.Equal(t, result, goModPath)
	}
}
