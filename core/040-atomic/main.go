package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	var counter int64
	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(i int) {
			//fmt.Println("GoRoutines: ", runtime.NumGoroutine())
			fmt.Println("counter: ", i+1, atomic.LoadInt64(&counter))
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter: ", i+1, atomic.LoadInt64(&counter))
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
