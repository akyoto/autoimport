package main

import (
	"os"
	"sync"

	"github.com/blitzprog/color"
)

// autoImport finds the correct import statements and writes them to the given Go source files.
func autoImport(files []string) {
	goPath := os.Getenv("GOPATH")
	goRoot := os.Getenv("GOROOT")

	if goRoot == "" {
		goRoot = "/usr/local/go"
	}

	if goPath == "" {
		home, err := os.UserHomeDir()

		if err != nil {
			panic(err)
		}

		goPath = home + "/go"
	}

	standardPackages := getPackagesInDirectory(goRoot)
	workspacePackages := getPackagesInDirectory(goPath)

	for name, packageList := range workspacePackages {
		packagesByName[name] = append(packagesByName[name], packageList...)
	}

	for name, packageList := range standardPackages {
		packagesByName[name] = append(packagesByName[name], packageList...)
	}

	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		path := file

		go func() {
			err := processFile(path)

			if err != nil {
				color.Red(err.Error())
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
