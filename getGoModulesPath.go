package autoimport

import (
	"os"
	"path"
)

// getGoModulesPath will attempt to find the path of the installed Go standard library.
func getGoModulesPath() string {
	goPath := os.Getenv("GOPATH")

	if goPath != "" {
		return path.Join(goPath, "pkg", "mod")
	}

	home, err := os.UserHomeDir()

	if err != nil {
		return ""
	}

	return pickExistingDirectory(
		path.Join(home, "go", "pkg", "mod"),
	)
}
