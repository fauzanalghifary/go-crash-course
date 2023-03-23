package main

import "fmt"

func main() {
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	fmt.Println(intMap)
	fmt.Println("-----")
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	delete(intMap, "three")
	fmt.Println(intMap)

	el, ok := intMap["four"]
	fmt.Println(el, ok)

	intMap["four"] = 44
	fmt.Println(intMap)
}
