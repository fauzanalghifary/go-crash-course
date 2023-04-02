package main

import "errors"

func main() {
	a := 12
	b := 6

	//if b != 0 {
	//	c := divideTwoNumbers(a, b)
	//
	//	if c == 2 {
	//		println("a divided by b is 2")
	//	}
	//}

	//if b != 0 && divideTwoNumbers(a, b) == 2 {
	//	println("a divided by b is 2")
	//}

	c, err := divideTwoNumbers(a, b)
	if err != nil {
		println(err.Error())
	} else {
		if c == 2 {
			println("a divided by b is 2")
		}
	}

}

func divideTwoNumbers(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
