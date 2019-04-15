package autoimport

import (
	"os"
	"path"
)

var packagesByName = map[string][]*Package{}

// Files finds the correct import statements and writes them to the given Go source files.
func Files(files []string) error {
	goPath := os.Getenv("GOPATH")
	goRoot := os.Getenv("GOROOT")

	if goRoot == "" {
		goRoot = "/usr/local/go"
	}

	if goPath == "" {
		home, err := os.UserHomeDir()

		if err != nil {
			return err
		}

		goPath = home + "/go"
	}

	standardPackages := GetPackagesInDirectory(path.Join(goRoot, "src"))
	packagesByName = standardPackages

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
