package autoimport

import (
	"github.com/blitzprog/autoimport/parser"
	"github.com/fatih/color"
)

// AutoImport helps you automatically add imports to your source files.
type AutoImport struct {
	index           PackageIndex
	moduleDirectory string
}

// New creates a new auto import.
func New(moduleDirectory string) *AutoImport {
	return &AutoImport{
		index:           PackageIndex{},
		moduleDirectory: moduleDirectory,
	}
}

// Source finds the correct import statements and returns code that includes import paths.
func (importer *AutoImport) Source(src []byte) []byte {
	identifiers := parser.PackageIdentifiers(src)
	println("Identifiers:")

	for id := range identifiers {
		color.Green(id)
	}

	return src
}
