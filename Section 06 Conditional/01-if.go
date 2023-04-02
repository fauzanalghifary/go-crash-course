package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(2)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ = reader.ReadString('\n')

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose rock")
	case PAPER:
		fmt.Println("Computer chose paper")
	case SCISSORS:
		fmt.Println("Computer chose scissors")
	default:
		fmt.Println("Computer chose something else")
	}

	if playerValue == computerValue {
		fmt.Println("It's a tie")
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				fmt.Println("You lose")
			} else {
				fmt.Println("You win")
			}
		case PAPER:
			if computerValue == SCISSORS {
				fmt.Println("You lose")
			} else {
				fmt.Println("You win")
			}
		case SCISSORS:
			if computerValue == ROCK {
				fmt.Println("You lose")
			} else {
				fmt.Println("You win")
			}
		default:
			fmt.Println("Invalid player choice")
		}

	}

}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
