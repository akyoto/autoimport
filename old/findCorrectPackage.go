package autoimport

func findCorrectPackage(packages []*Package) *Package {
	// // Choose the shortest path
	// sort.Slice(packages, func(i, j int) bool {
	// 	return len(packages[i].RealPath) < len(packages[j].RealPath)
	// })
	return packages[0]
}
