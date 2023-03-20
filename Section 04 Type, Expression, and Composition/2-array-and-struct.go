package main

import "fmt"

func main() {
	// Array
	var myArray [5]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5
	fmt.Println(myArray)

	//	Struct
	type Car struct {
		Brand         string
		Year          int
		NumberOfTires int
	}

	var myCar Car
	myCar.Brand = "Toyota"
	myCar.Year = 2019
	myCar.NumberOfTires = 4

	yourCar := Car{
		Brand:         "Honda",
		Year:          2018,
		NumberOfTires: 4,
	}
	fmt.Println(myCar, yourCar)
}
