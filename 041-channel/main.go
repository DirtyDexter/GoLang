package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() {
		fmt.Println("before receiving")
		fmt.Println(<-c)
		fmt.Println("after receiving")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 5)
		c <- 42
		wg.Done()
	}()

	wg.Wait()

	d := make(chan int, 1)
	d <- 43
	fmt.Println(<-d)
}
