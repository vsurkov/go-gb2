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

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(sum *int64) {
			defer wg.Done()
			atomic.AddInt64(sum, 1)
		}(&sum)
	}

	wg.Wait()
	fmt.Printf("Done, all Goroutines %d was ended", sum)
}
