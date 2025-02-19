package main

import (
	"fmt"
	"sync"
)

func manageTicket(ticketChanel <-chan int, doneChanel <-chan bool, totalTicket *int) {
	for {
		select {
		case userId := <-ticketChanel:
			if *totalTicket > 0 {
				fmt.Printf("User %d can purchase ticket\n", userId)
				*totalTicket--
			} else {
				fmt.Printf("No ticket available\n")
			}
		case <-doneChanel:
			fmt.Println("DONE")
		}
	}
}

func buyTicket(ticketChanel chan<- int, userId int, wg *sync.WaitGroup) {
	defer wg.Done()

	ticketChanel <- userId

}

func main() {
	totalTicket := 5
	totalUser := 10

	var wg sync.WaitGroup

	ticketChanel := make(chan int)
	doneChanel := make(chan bool)

	go manageTicket(ticketChanel, doneChanel, &totalTicket)

	for userId := 0; userId <= totalUser; userId++ {
		wg.Add(1)
		go buyTicket(ticketChanel, userId, &wg)
	}

	wg.Wait()

	close(ticketChanel)

}