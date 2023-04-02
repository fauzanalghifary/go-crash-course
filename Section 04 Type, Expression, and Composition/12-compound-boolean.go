package main

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	jack := Employee{
		Name:     "Jack",
		Age:      23,
		Salary:   3000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill",
		Age:      23,
		Salary:   4000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.Age >= 23 {
			println(x.Name, "is older than 23")
		}

		if x.Salary >= 3000 && x.Age > 30 {
			println(x.Name, "makes more than 3000 and is older than 30")
		}

		if x.FullTime || x.Age > 30 {
			println(x.Name, "is full time or older than 30")
		}

		if (x.FullTime || x.Age > 30) && x.Salary > 3000 {
			println(x.Name, "is full time or older than 30 and makes more than 3000")
		}
	}

}
