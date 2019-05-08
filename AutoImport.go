package autoimport

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/akyoto/autoimport/parser"
	"github.com/akyoto/color"
)

var packageStatement = []byte("package ")

// AutoImport helps you automatically add imports to your source files.
type AutoImport struct {
	index           PackageIndex
	moduleDirectory string
}

// New creates a new auto import.
func New(moduleDirectory string) *AutoImport {
	standardPackagesPath := getStandardPackagesPath()
	standardPackages := GetPackagesInDirectory(standardPackagesPath, standardPackagesPath)
	goModPath := findGoMod(moduleDirectory)
	dependencies, moduleImportPath, err := readGoMod(goModPath)

	if err != nil {
		panic(err)
	}

	// Find where modules are cached
	goModulesPath := getGoModulesPath()

	for _, dep := range dependencies {
		directoryName, err := encodePkgModPath(dep.ImportPath)

		if err != nil {
			color.Red(err.Error())
			continue
		}

		directoryName = fmt.Sprintf("%s@%s", directoryName, dep.Version)
		packageLocation := path.Join(goModulesPath, directoryName)
		importedPackages := GetPackagesInDirectory(packageLocation, goModulesPath)
		merge(standardPackages, importedPackages)
	}

	// Local packages
	innerPackages := GetPackagesInDirectory(moduleDirectory, moduleDirectory)

	for _, packageList := range innerPackages {
		for i := range packageList {
			packageList[i].ImportPath = fmt.Sprintf("%s/%s", moduleImportPath, packageList[i].ImportPath)
		}
	}

	merge(standardPackages, innerPackages)

	return &AutoImport{
		index:           standardPackages,
		moduleDirectory: moduleDirectory,
	}
}

// Imports returns the import paths used in the given source file.
func (importer *AutoImport) Imports(src []byte) ([]string, error) {
	imports := make([]string, 0, 16)
	identifiers, err := parser.PackageIdentifiers(src)

	if err != nil {
		return nil, err
	}

	for id := range identifiers {
		possiblePackages := importer.index[id]

		if len(possiblePackages) == 0 {
			continue
		}

		pkg := findCorrectPackage(possiblePackages)
		imports = append(imports, pkg.ImportPath)
	}

	// Sort by file system depth
	sort.Slice(imports, func(a int, b int) bool {
		countA := strings.Count(imports[a], "/")
		countB := strings.Count(imports[b], "/")

		if countA == countB {
			return imports[a] < imports[b]
		}

		return countA < countB
	})

	return imports, nil
}

// Source finds the correct import statements and returns code that includes import paths.
func (importer *AutoImport) Source(src []byte) ([]byte, error) {
	imports, err := importer.Imports(src)

	if err != nil {
		return src, err
	}

	// for _, importPath := range imports {
	// 	fmt.Printf("%s\n", color.GreenString(importPath))
	// }

	importCommand := ""

	if len(imports) > 0 {
		importCommand = fmt.Sprintf("\nimport (\n\t\"%s\"\n)\n", strings.Join(imports, "\"\n\t\""))
	}

	// Find package definition
	packagePos := bytes.Index(src, packageStatement)

	if packagePos == -1 {
		return src, errors.New("Package definition missing")
	}

	seekPos := 0

	for i := packagePos; i < len(src); i++ {
		if src[i] == '\n' {
			seekPos = i + 1
			break
		}
	}

	// Contains only a single package statement
	if seekPos == 0 {
		return src, nil
	}

	// Insert imports
	endOfFile := append([]byte(importCommand), src[seekPos:]...)
	newSource := append(src[:seekPos], endOfFile...)

	return newSource, nil
}
