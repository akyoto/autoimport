package autoimport

import (
	"os"
	"path"
)

// findStandardPackagesPath will attempt to find the path of the installed Go standard library.
func findStandardPackagesPath() string {
	goRoot := os.Getenv("GOROOT")

	if goRoot != "" {
		return path.Join(goRoot, "src")
	}

	return pickExistingDirectory(
		"/usr/local/go/src",
		"/usr/lib/go/src/",
	)
}
