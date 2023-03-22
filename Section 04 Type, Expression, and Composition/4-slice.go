package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "bird")
	animals = append(animals, "fish")

	fmt.Println(animals)

	for _, animal := range animals {
		fmt.Println(animal)
	}

	fmt.Println(animals[1:3])
	fmt.Println(animals[1])
	fmt.Println(len(animals))
	fmt.Println(sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println(sort.StringsAreSorted(animals))
	fmt.Println(animals)
	fmt.Println("------")
	deleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
