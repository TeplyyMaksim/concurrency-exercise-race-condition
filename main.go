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

	wg.Add(incrementsAmount)
	for i := 0; i < incrementsAmount; i++  {
		go func () {
			sumCopy := sum
			runtime.Gosched()
			sumCopy += 1
			sum = sumCopy
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines amount in the end:", runtime.NumGoroutine())
	fmt.Println(sum)
}
