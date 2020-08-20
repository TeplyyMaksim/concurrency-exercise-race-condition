package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main()  {
	var sum uint64
	incrementsAmount := 100

	wg.Add(incrementsAmount)
	for i := 0; i < incrementsAmount; i++  {
		go func () {
			atomic.AddUint64(&sum, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines amount in the end:", runtime.NumGoroutine())
	fmt.Println(sum)
}
