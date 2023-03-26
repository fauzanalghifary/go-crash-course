package main

import "fmt"

func main() {
	age := 10
	name := "John"
	rightHanded := true

	fmt.Printf("%s is %d years old and is %t rightHanded", name, age, rightHanded)

	ageInTenYears := age + 10
	isAgeTeenager := age >= 13 && age <= 19

	fmt.Println(ageInTenYears, isAgeTeenager)
}
