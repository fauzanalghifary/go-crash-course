package main

import "fmt"

func main() {

	//	multiplication and division
	a := 10 * 5 / 2
	b := 10 * (5 / 2)
	fmt.Println(a, b)

	//	integer division
	unclear := 10 / 4
	clear := 10.0 / 4.0
	fmt.Println(unclear, clear)

	//	parentheses
	c := (10 + 2) * 3 / 2
	fmt.Println(c)

	//	addition and subtraction
	d := 10 + 5 - 2
	e := 10 + (5 - 2)
	fmt.Println(d, e)
}
