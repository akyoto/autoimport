package autoimport

import (
	"os"
)

// pickExistingDirectory returns the first existing directory.
func pickExistingDirectory(directories ...string) string {
	for _, directory := range directories {
		stat, err := os.Stat(directory)

		if err == nil && stat.IsDir() {
			return directory
		}
	}

	return ""
}
