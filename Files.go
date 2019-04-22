package autoimport

var packagesByName = map[string][]*Package{}

// Files finds the correct import statements and writes them to the given Go source files.
func Files(files []string) error {
	if len(files) == 0 {
		return nil
	}

	// Scan the standard library
	standardPackagesPath := findStandardPackagesPath()
	standardPackages := GetPackagesInDirectory(standardPackagesPath)
	packagesByName = standardPackages

	// Find go.mod file
	goModPath := FindGoMod(files[0])

	// Scan go.mod required dependencies
	println(goModPath)
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
