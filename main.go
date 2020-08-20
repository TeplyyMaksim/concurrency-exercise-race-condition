package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main()  {
	sum := 0
	incrementsAmount := 100
	var mutex = &sync.Mutex{}

	wg.Add(incrementsAmount)
	for i := 0; i < incrementsAmount; i++  {
		go func () {
			mutex.Lock()
			sumCopy := sum
			runtime.Gosched()
			sumCopy += 1
			sum = sumCopy
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines amount in the end:", runtime.NumGoroutine())
	fmt.Println(sum)
}
