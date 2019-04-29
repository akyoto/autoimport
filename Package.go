package autoimport

// Package represents a single Go package.
type Package struct {
	Name          string // Actual package name as declared by the package keyword
	RealPath      string // Real file path
	ImportPath    string // Import path used in Go source files
	DirectoryName string // Base name of the directory
	IsModuleRoot  bool   // Only root packages have go.mod files
}
