package autoimport

import (
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/akyoto/color"
)

// getPackagesInDirectory returns a map of package names mapped to packages.
func getPackagesInDirectory(srcPath string, root string) PackageIndex {
	// fmt.Println("Scanning", srcPath)
	packages := []*Package{}
	packageByPath := map[string]*Package{}
	packagesByName := map[string][]*Package{}
	packagePrefix := "\npackage "

	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	// onDirectory
	onDirectory := func(path string) error {
		baseName := filepath.Base(path)

		if strings.HasPrefix(baseName, ".") || strings.HasPrefix(baseName, "_") || baseName == "vendor" || baseName == "builtin" || baseName == "internal" || baseName == "testdata" {
			return filepath.SkipDir
		}

		packageRealPath := strings.TrimPrefix(path, root)
		packageName := baseName

		if packageName == "." {
			return nil
		}

		// Remove version number from import path
		packageImportPath := packageRealPath
		atPosition := strings.Index(packageImportPath, "@")

		if atPosition != -1 {
			slashPosition := strings.Index(packageImportPath[atPosition:], "/")

			if slashPosition == -1 {
				packageImportPath = packageImportPath[:atPosition]
			} else {
				packageImportPath = packageImportPath[:atPosition] + packageImportPath[atPosition+slashPosition:]
			}
		}

		pkg := &Package{
			DirectoryName: packageName,
			RealPath:      packageRealPath,
			ImportPath:    packageImportPath,
		}

		packages = append(packages, pkg)
		packageByPath[pkg.RealPath] = pkg

		// fmt.Println(color.GreenString(pkg.DirectoryName), pkg.RealPath)
		return nil
	}

	// onFile
	onFile := func(path string) error {
		// Ignore files that are not Go source code
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		packageRealPath := filepath.Dir(path)
		packageRealPath = strings.TrimPrefix(packageRealPath, root)
		// fmt.Println("Go file in", packageRealPath, filepath.Base(path))

		pkg, exists := packageByPath[packageRealPath]

		if !exists {
			return nil
		}

		if pkg.Name != "" {
			return nil
		}

		file, err := os.Open(path)

		if err != nil {
			return err
		}

		defer file.Close()
		buffer := make([]byte, 1+4096)
		buffer[0] = '\n'
		_, err = file.Read(buffer[1:])

		if err != nil {
			return err
		}

		packageLine := string(buffer)
		prefixPos := strings.Index(packageLine, packagePrefix)

		if prefixPos == -1 {
			return nil
		}

		packageLine = packageLine[prefixPos+len(packagePrefix):]

		for index, char := range packageLine {
			if unicode.IsSpace(char) {
				pkg.Name = packageLine[:index]

				if pkg.Name == "main" {
					return nil
				}

				packagesByName[pkg.Name] = append(packagesByName[pkg.Name], pkg)
				// fmt.Printf("Package %s in directory %s\n", color.GreenString(pkg.Name), color.YellowString(pkg.DirectoryName))
				return nil
			}
		}

		return nil
	}

	// Traverse directory
	filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			color.Red(err.Error())
			return nil
		}

		if info == nil {
			color.Red("Invalid file path: %s", srcPath)
			return nil
		}

		if !info.IsDir() {
			return onFile(path)
		}

		return onDirectory(path)
	})

	return packagesByName
}
