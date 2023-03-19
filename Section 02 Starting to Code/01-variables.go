package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready"

func main() {
	//	seed the random number generator
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var thirdNumber = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - thirdNumber

	playTheGame(firstNumber, secondNumber, thirdNumber, answer)

}

func playTheGame(firstNumber int, secondNumber int, thirdNumber int, answer int) {
	reader := bufio.NewReader(os.Stdin)
	//	display a welcome / instructions message
	fmt.Println("Guess the number game")
	fmt.Println("---------------------------")
	fmt.Println("")

	fmt.Println("think of a number between 1 and 10", prompt)
	var _, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	//	take them through the game
	fmt.Println("multiply your number by", firstNumber, prompt)
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("multiply the result by", secondNumber, prompt)
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("divide the result by the number you originally thought of", prompt)
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("now subtract", thirdNumber, prompt)
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	// give them the answer
	fmt.Println("The answer is", answer)
}
