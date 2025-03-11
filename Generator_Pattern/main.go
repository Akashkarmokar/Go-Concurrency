package main

import "fmt"

func generateNumber(num int) chan int {

	// Create a chanel
	out := make(chan int)

	go func() {
		defer close(out)

		for i := 1; i <= num; i++ {
			out <- i
		}
	}()

	// return chanel
	return out
}

func main() {

	ourChanel := generateNumber(4)

	// recieve

	for num := range ourChanel {
		fmt.Println("Reviews values: ", num)
	}
}
