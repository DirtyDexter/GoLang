package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)

	c2 := make(chan int)
	quit2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit2 <- 0
	}()
	foo(c2, quit2)

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// receive channel
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok bit1", i, ok)
				return
			} else {
				fmt.Println("from comma ok bit2", i, ok)
			}
		}
	}
}

func foo(c2, quit2 chan int) {
	x := 3
	for {
		select {
		case c2 <- x:
			x += x
		case <-quit2:
			fmt.Println("about to exit2")
			return
		}
	}
}
