package main

import (
	"fmt"
	"sync"
)

//2. Реализуйте функцию для разблокировки мьютекса с помощью defer

func main() {
	const (
		count = 1000
	)

	var sum = 0
	var wg sync.WaitGroup
	var mux sync.Mutex

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(sum *int) {
			defer func() {
				mux.Unlock()
				wg.Done()
			}()
			mux.Lock()
			*sum += 1
		}(&sum)
	}

	wg.Wait()
	fmt.Printf("Done, all Goroutines %d was ended", sum)
}
