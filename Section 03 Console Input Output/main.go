package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main2() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Espresso"
	coffees[3] = "Latte"
	coffees[4] = "Mocha"
	coffees[5] = "Tea"
	coffees[6] = "Quit"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Espresso")
	fmt.Println("3 - Latte")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Tea")
	fmt.Println("6 - Quit")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You choose %s", coffees[i]))

	}

	fmt.Println("Program terminated")
}
