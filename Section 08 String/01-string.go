package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello World"
	fmt.Println(name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}

	fmt.Println()
	fmt.Println("Index\tRune\tString")
	for i, v := range name {
		fmt.Println(i, "\t", v, "\t", string(v))
	}

	h := "Hello"
	w := "World"

	//	using + (not very efficient)
	fmt.Println(h + " " + w)
	myString := h + w
	fmt.Println(myString)

	// using fmt sprintF (more efficient)
	myString = fmt.Sprintf("%s %s", h, w)
	fmt.Println(myString)

	//	using stringbuilder (most efficient)
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(" ")
	sb.WriteString(w)
	fmt.Println(sb.String())

	//	substring
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(name[0:3])
	fmt.Println(name[3:6])

}
