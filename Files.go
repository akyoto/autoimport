package autoimport

import (
	"fmt"
	"path"
)

var packagesByName = PackageIndex{}

// Files finds the correct import statements and writes them to the given Go source files.
func Files(files []string) error {
	if len(files) == 0 {
		return nil
	}

	// Scan the standard library
	standardPackagesPath := findStandardPackagesPath()
	standardPackages := GetPackagesInDirectory(standardPackagesPath, standardPackagesPath)
	packagesByName = standardPackages

	// Find go.mod file
	goModPath := FindGoMod(files[0])

	// Scan go.mod required dependencies
	dependencies, err := GetDirectDependencies(goModPath)

	if err != nil {
		return err
	}

	// Find where modules are cached
	goModulesPath := findGoModulesPath()

	for _, dep := range dependencies {
		directoryName := fmt.Sprintf("%s@%s", dep.ImportPath, dep.Version)
		packageLocation := path.Join(goModulesPath, directoryName)
		importedPackages := GetPackagesInDirectory(packageLocation, goModulesPath)

		for name, packageList := range importedPackages {
			packagesByName[name] = append(packageList, packagesByName[name]...)
		}
	}

	printPackagesByName(packagesByName)

	// ...

	// wg := sync.WaitGroup{}

	// for _, file := range files {
	// 	wg.Add(1)
	// 	path := file

	// 	go func() {
	// 		err := processFile(path)

	// 		if err != nil {
	// 			color.Red(err.Error())
	// 		}

	// 		wg.Done()
	// 	}()
	// }

	// wg.Wait()
	return nil
}
