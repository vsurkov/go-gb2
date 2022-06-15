package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/vsurkov/lgr/v2"
)

//С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из которых увеличивает число на 1.
//Дождаться завершения всех горутин и убедиться, что при каждом запуске программы итоговое число равно 1000.
func main() {
	const (
		count    = 10000
		buffSize = 10
	)

	var workers = make(chan struct{}, buffSize)
	var result = make(chan int64)
	var sum int64

	// горутина читает значения из канала
	var wgl sync.WaitGroup
	wgl.Add(1)

	go func() {
		defer wgl.Done()
		for val := range result {
			atomic.AddInt64(&sum, val)
		}
	}()

	// горутины пишут значения в канал и добавляются в WaitGroup
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		workers <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			result <- 1
			<-workers
		}()
	}
	wg.Wait()
	close(result)
	wgl.Wait()
	lgr.Logger(fmt.Sprintf("result: %d", sum))
}
