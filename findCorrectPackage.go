package main

func findCorrectPackage(packages []*Package) *Package {
	// // Choose the shortest path
	// sort.Slice(packages, func(i, j int) bool {
	// 	return len(packages[i].Path) < len(packages[j].Path)
	// })
	return packages[0]
}
