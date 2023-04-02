package main

import (
	"math/rand"
	"time"
)

func main() {
	// while loop
	rand.Seed(time.Now().UnixNano())
	i := 1000
	for i > 100 {
		i = rand.Intn(1000) + 1
		println(i)
	}
}
