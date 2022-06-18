package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

//2. Написать многопоточную программу, в которой будет использоваться явный вызов планировщика. Выполните трассировку программы

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var wgl sync.WaitGroup
	var wgw sync.WaitGroup

	var count = 1000
	var listeners = 10

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
			data <- w + 1
			//явный вызов планировщика
			runtime.Gosched()
		}(w)
	}

	wgw.Wait()
	close(data)
	wgl.Wait()

	fmt.Println(sum)
}
