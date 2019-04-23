package autoimport

import (
	"io"
	"os"
	"strings"
)

// readGoMod will attempt to read all direct dependencies in the go.mod file.
// It returns a list of all dependencies used, the module import path and potentially an error.
func readGoMod(goModPath string) ([]*Dependency, string, error) {
	file, err := os.Open(goModPath)

	if err != nil {
		return nil, "", err
	}

	defer file.Close()
	buffer := make([]byte, 4096)
	lastNewline := 0
	remaining := ""
	moduleImportPath := ""
	dependencies := []*Dependency{}

	for {
		n, err := file.Read(buffer)

		for i := 0; i < n; i++ {
			if buffer[i] == '\n' {
				line := remaining + string(buffer[lastNewline:i])
				line = strings.TrimSpace(line)
				lastNewline = i
				remaining = ""

				if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "replace ") {
					continue
				}

				if strings.HasPrefix(line, "module ") {
					moduleImportPath = line[len("module "):]
					continue
				}

				// Ignore indirect dependencies
				if strings.HasSuffix(line, " // indirect") {
					continue
				}

				// Only take lines that have 2 slashes
				if strings.Count(line, "/") < 2 {
					continue
				}

				// Remove require keyword
				line = strings.TrimPrefix(line, "require ")

				// Remove version number
				space := strings.Index(line, " ")

				// Add to list
				dependencies = append(dependencies, &Dependency{
					ImportPath: line[:space],
					Version:    strings.TrimSpace(line[space:]),
				})
			}
		}

		remaining = string(buffer[lastNewline:n])

		if err == io.EOF {
			break
		}

		lastNewline = 0
	}

	return dependencies, moduleImportPath, nil
}
