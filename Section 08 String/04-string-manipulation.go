package main

import (
	"fmt"
	"strings"
)

func main() {
	newString := "Go is a beautiful language! Go for it!"

	if strings.Contains(newString, "Go") {
		//newString = strings.Replace(newString, "Go", "Golang", 1)
		//newString = strings.Replace(newString, "Go", "Golang", -1)
		newString = strings.ReplaceAll(newString, "Go", "Golang")
	}

	println(newString)

	//	String comparison
	println(strings.Compare("A", "B"))
	println(strings.Compare("B", "B"))

	badEmail := " me@here.com"
	println(strings.TrimSpace(badEmail))

	str := "alpha alpha alpha alpha"
	str = replaceNth(str, "alpha", "beta", 2)
	fmt.Println(str)
}

func replaceNth(s, old, new string, n int) string {
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			// Not found
			break
		}
		// Found at index i
		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}

		i = i + len(old)
	}

	return s
}
