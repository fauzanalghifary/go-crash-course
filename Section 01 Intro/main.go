package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print(">> ")
		userInput, _ := reader.ReadString('\n')
		//userInput = userInput[:len(userInput)-1]
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		}

		response := doctor.Response(userInput)
		fmt.Println(response)
	}

}
