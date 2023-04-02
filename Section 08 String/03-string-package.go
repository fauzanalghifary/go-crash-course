package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := []string{
		"Learn Go for Beginners",
		"Learn Java for Beginners",
		"Learn Python for Beginners",
		"Learn C++ for Beginners",
	}

	//	contains
	for _, course := range courses {
		if strings.Contains(course, "Go") {
			fmt.Println("Go is found in", course, "and the index is", strings.Index(course, "Go"))
		}
	}

	newString := "Go is a great programming language. Go for it!"
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "!"))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "Fish"))
	fmt.Println(strings.Index(newString, "Go"))
	fmt.Println(strings.Index(newString, "Python"))
	fmt.Println(strings.LastIndex(newString, "Go"))

}
