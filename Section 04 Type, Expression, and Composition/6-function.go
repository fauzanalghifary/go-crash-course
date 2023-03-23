package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Speak() {
	fmt.Println(a.Name, a.Sound)
}

func main() {
	fmt.Println(addTwoNumbers(2, 3))
	fmt.Println(sumManyNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	dog := Animal{
		Name:         "Dog",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}
	dog.Speak()
}

func addTwoNumbers(x int, y int) int {
	return x + y
}

func sumManyNumbers(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
