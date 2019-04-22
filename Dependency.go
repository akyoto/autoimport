package autoimport

// Dependency represents a single package dependency in a go.mod file.
type Dependency struct {
	ImportPath string
	Version    string
}
