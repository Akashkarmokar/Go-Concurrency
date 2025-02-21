package main

import (
	"fmt"
	"sync"
)

func generator(maxNumber int, intChan chan<- int, parentWg *sync.WaitGroup) {
	defer parentWg.Done()
	for i := 1; i <= maxNumber; i++ {
		intChan <- i
	}
	close(intChan)
}

func main() {
	intChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go generator(5, intChan, &wg)

	for value := range intChan {
		fmt.Println("Values: ", value)
	}
	wg.Wait()

}
