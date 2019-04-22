package autoimport

import (
	"fmt"

	"github.com/blitzprog/color"
)

func printPackagesByName(packagesByName map[string][]*Package) {
	// Output
	for name, packageList := range packagesByName {
		fmt.Println(color.GreenString(name))

		for _, pkg := range packageList {
			fmt.Printf(" - %s\n", pkg.ImportPath)
		}
	}
}
