package main

import "fmt"

type Package struct {
	name string
	deps []Package
}

func main() {
	/*
		"A" : ["D", "B"]
		"B" : ["C", "E"]
		"C" : ["D", "E"]
	*/

	D := addPackage("D")
	E := addPackage("E")
	C := addPackage("C", D, E)
	B := addPackage("B", C, E)
	A := addPackage("A", D, B)

	packages := getDeps(A)
	fmt.Printf("Package: %v depends on packages: %v\n", A.name, packages)

}

func addPackage(name string, deps ...Package) Package {
	newPackage := Package{
		name: name,
		deps: deps,
	}
	return newPackage
}

func getDeps(p Package) []string {
	var result []string
	var names = make(map[string]bool)
	ret := helper(p, result, names)
	return ret
}

func helper(dep Package, result []string, names map[string]bool) []string {
	for _, x := range dep.deps {
		if !names[x.name] {
			names[x.name] = true
			result = append(result, x.name)
			if x.deps != nil {
				return helper(x, result, names)
			}
		}
	}
	return result
}
