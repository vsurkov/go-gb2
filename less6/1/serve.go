package main

import (
	"sync"
)

func serve(count int, listeners int) int {
	var wgl sync.WaitGroup
	var wgw sync.WaitGroup

	var wp = make(chan struct{}, listeners)
	var data = make(chan int, 3)

	var sum int
	var mux sync.Mutex

	for l := 0; l < listeners; l++ {
		wgl.Add(1)
		go func() {
			defer wgl.Done()
			for value := range data {
				mux.Lock()
				sum += value
				mux.Unlock()
			}
		}()
	}

	for w := 0; w < count; w++ {
		wgw.Add(1)
		wp <- struct{}{}

		go func(w int) {
			defer wgw.Done()
			<-wp
			data <- w
		}(w)
	}

	wgw.Wait()
	close(data)
	wgl.Wait()

	return sum
}
