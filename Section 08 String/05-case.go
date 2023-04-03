package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "This is a clear example of why we search in one case only"
	searchString := strings.ToLower(myString)
	if strings.Contains(searchString, "this") {
		println("Found it!")
	} else {
		println("Not found")
	}

	fmt.Println(strings.ToLower("Hello, World!"))
	fmt.Println(strings.ToUpper("Hello, World!"))
	fmt.Println(strings.Title("hello, world!"))
}
