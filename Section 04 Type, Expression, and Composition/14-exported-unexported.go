package main

import (
	"fmt"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Mary", LastName: "Taylor", Salary: 40000, FullTime: true},
}

func main() {
	myStaff := staff.Office{AllStaff: employees}

	fmt.Println(myStaff.All())
	staff.OverPaidLimit = 50000
	fmt.Println(myStaff.Overpaid())
	fmt.Println(myStaff.Underpaid())
}
