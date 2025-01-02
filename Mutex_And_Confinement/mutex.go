package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func processTicket(wg *sync.WaitGroup, userId int, reamainingTicket *int) {
	defer wg.Done()

	mutex.Lock()
	if *reamainingTicket > 0 {
		*reamainingTicket--
		fmt.Printf("User %d purchased a ticket and reamaining %d\n", userId, *reamainingTicket)
	} else {
		fmt.Printf("User %d not found any ticket\n", userId)
	}
	mutex.Unlock()
}
func main() {
	var totalTicket int = 5

	var wg sync.WaitGroup
	for userId := 0; userId < 10; userId++ {
		wg.Add(1)
		go processTicket(&wg, userId, &totalTicket)
	}
	wg.Wait()
}
