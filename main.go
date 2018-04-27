package main

import (
	"flag"
	"os"
	"sync"

	"github.com/fatih/color"
)

var packagesByName = map[string][]*Package{}

func main() {
	flag.Parse()
	files := flag.Args()
	goPath := os.Getenv("GOPATH")
	goRoot := os.Getenv("GOROOT")

	if goRoot == "" {
		goRoot = "/usr/local/go"
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
