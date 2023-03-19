package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "One"

func main() {
	//var one = "this is a block scoped variable"
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println(newString)

	//privateString := packageone.privateVar
	//fmt.Println(privateString)

	packageone.PublicFunc()
}

func myFunc() {
	fmt.Println(one)
}
