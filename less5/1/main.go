package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//1. Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех

func main() {
	const (
		count = 1000
	)

	var sum int64 = 0
	var wg sync.WaitGroup
	var ch = make(chan struct{}, count)

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func(sum *int64) {
			defer wg.Done()
			atomic.AddInt64(sum, 1)
			ch <- struct{}{}
		}(&sum)
	}
	wg.Wait()
	close(ch)

	i := 0
	for range ch {
		i += 1
	}
	fmt.Printf("Done, all Goroutines %d was ended and %d calculated", sum, i)
}
