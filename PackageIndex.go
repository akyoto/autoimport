package autoimport

// PackageIndex is a map that maps a package name to possible import paths.
type PackageIndex = map[string][]*Package

// merge merges `b` into `a` while giving the packages in `b` higher priority.
func merge(a PackageIndex, b PackageIndex) {
	for name, packageList := range b {
		a[name] = append(packageList, a[name]...)
	}
}
