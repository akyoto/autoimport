package autoimport

// PackageIndex is a map that maps a package name to possible import paths.
type PackageIndex = map[string][]*Package
