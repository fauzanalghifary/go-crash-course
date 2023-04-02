package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func ListenForLog(ch chan string) {
	for {
		msg := <-ch
		log.Println(msg)
	}
}

func main() {
	//	 read input from the user 5 times and write it to a log
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go ListenForLog(ch)

	fmt.Println("Enter a string: ")

	for i := 0; i < 5; i++ {
		fmt.Println("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
