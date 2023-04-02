package main

import "fmt"

func main() {
	courseName := "Learn Go for Beginners Crash Course"

	fmt.Println(courseName[0])         // 76
	fmt.Println(string(courseName[0])) // L

	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

	fmt.Println("Length of courseName: ", len(courseName)) // 30

	var mySlice []string
	mySlice = append(mySlice, "Hello")
	mySlice = append(mySlice, "World")

	fmt.Println("mySlice has a length of: ", len(mySlice)) // 2
	fmt.Println("the last element of mySlice is: ", mySlice[1])

}
