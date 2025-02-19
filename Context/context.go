package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func func1(ctx context.Context, parentWg *sync.WaitGroup, stream <-chan interface{}) {
	defer parentWg.Done()

	var wg sync.WaitGroup

	doWork := func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Func 1 cancel by parent fun")
				return
			case d, ok := <-stream:
				if !ok {
					fmt.Println("Channel Closed of func1")
					return
				}
				fmt.Println(d)
			}
		}
	}

	newCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doWork(newCtx)
	}

	wg.Wait()
}

func genericFunc(ctx context.Context, wg *sync.WaitGroup, stream <-chan interface{}) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Channel Closed from parent ......")
			return

		case d, ok := <-stream:
			if !ok {
				fmt.Println("Chanel Closed")
				return
			}
			fmt.Println("Data from Generic: ", d)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	generator := func(dataItem string, stream chan interface{}) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker cancel by parent")
				return
			case stream <- dataItem:
			}
		}
	}

	infiniteApples := make(chan interface{})
	go generator("apple", infiniteApples)

	infiniteOrange := make(chan interface{})
	go generator("orange", infiniteOrange)

	infinitePeaches := make(chan interface{})
	go generator("peache", infinitePeaches)

	wg.Add(1)
	go func1(ctx, &wg, infiniteApples)

	func2 := genericFunc
	func3 := genericFunc

	wg.Add(1)
	go func2(ctx, &wg, infiniteOrange)

	wg.Add(1)
	go func3(ctx, &wg, infinitePeaches)

	wg.Wait()

}
