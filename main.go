package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
)

// Package ...
type Package struct {
	DirectoryName string
	Name          string
	Path          string
}

func main() {
	start := time.Now()
	goPath := os.Getenv("GOPATH")
	srcPath := path.Join(goPath, "src") + "/"
	packages := []*Package{}
	packageByPath := map[string]*Package{}
	packagesByName := map[string][]*Package{}
	packagePrefix := "package "

	filepath.Walk(goPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// Ignore files that are not Go source code
			if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return nil
			}

			packagePath := filepath.Dir(path)
			packagePath = strings.TrimPrefix(packagePath, srcPath)
			// fmt.Println("Go file in", packagePath, filepath.Base(path))

			pkg, exists := packageByPath[packagePath]

			if !exists {
				return nil
			}

			if pkg.Name == "" {
				file, err := os.Open(path)

				if err != nil {
					return err
				}

				defer file.Close()
				buffer := make([]byte, 256, 256)
				_, err = file.Read(buffer)

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
						packagesByName[pkg.Name] = append(packagesByName[pkg.Name], pkg)
						// fmt.Println(color.YellowString("CONFIRMED"), pkg.Name, pkg.DirectoryName)
						break
					}
				}
			}

			return nil
		}

		baseName := filepath.Base(path)

		if strings.HasPrefix(baseName, ".") || strings.HasPrefix(baseName, "_") {
			return filepath.SkipDir
		}

		packagePath := strings.TrimPrefix(path, srcPath)
		packageName := filepath.Base(packagePath)

		pkg := &Package{
			DirectoryName: packageName,
			Path:          packagePath,
		}

		packages = append(packages, pkg)
		packageByPath[pkg.Path] = pkg

		// fmt.Println(color.GreenString(pkg.DirectoryName), pkg.Path)

		if err != nil {
			color.Red(err.Error())
		}

		return err
	})

	// for name, packageList := range packagesByName {
	// 	fmt.Println(color.GreenString(name))

	// 	for _, pkg := range packageList {
	// 		fmt.Println(pkg.Path)
	// 	}
	// }

	fmt.Println(len(packagesByName), "packages", time.Since(start))
}
