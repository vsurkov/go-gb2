package main

import (
	"fmt"
	"github.com/vsurkov/lgr/v2"
	"sync"
	"sync/atomic"
	"time"
)

//С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из которых увеличивает число на 1.
//Дождаться завершения всех горутин и убедиться, что при каждом запуске программы итоговое число равно 1000.
func main() {
	const (
		count    = 100000
		buffSize = 10
	)

	var wg sync.WaitGroup
	var workers = make(chan struct{}, buffSize)
	var result = make(chan int64)
	var sum int64

	lgr.Logger(fmt.Sprintf("init start with buffSize=%v, count=%v", buffSize, count))

	// горутина читает значения из канала
	go func() {
		for val := range result {
			atomic.AddInt64(&sum, val)
		}
	}()

	// горутины пишут значения в канал и добавляются в WaitGroup
	for i := 0; i < count; i++ {
		workers <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				time.Sleep(time.Millisecond)
				wg.Done()
			}()
			result <- 1
			<-workers
		}()
	}

	wg.Wait()
	lgr.Logger(fmt.Sprintf("result: %d", sum))
}
