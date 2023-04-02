package main

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break
			}
			println("i = ", i, "j = ", j)
		}
	}
}
