package autoimport

import (
	"os"
	"path"
	"path/filepath"
)

// findGoMod will attempt to find the absolute path of the go.mod file.
func findGoMod(file string) string {
	directory, err := filepath.Abs(path.Clean(file))

	if err != nil {
		return ""
	}

	for {
		goMod := path.Join(directory, "go.mod")
		stat, err := os.Stat(goMod)

		if !os.IsNotExist(err) && stat != nil && !stat.IsDir() {
			return goMod
		}

		if directory == "/" {
			return ""
		}

		// Go up one level
		directory = path.Clean(path.Join(directory, ".."))
	}
}
