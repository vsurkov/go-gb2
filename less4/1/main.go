package main

import (
	"fmt"
	"github.com/vsurkov/lgr/v2"
	"sync"
	"time"
)

//С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из которых увеличивает число на 1.
//Дождаться завершения всех горутин и убедиться, что при каждом запуске программы итоговое число равно 1000.
func main() {
	const (
		count    = 1000
		buffSize = 10
	)

	var wg sync.WaitGroup
	var workers = make(chan struct{}, count)
	var result = make(chan int)
	var sum int

	lgr.Logger(fmt.Sprintf("init start with buffSize=%v, count=%v", buffSize, count))

	// горутина читает значения из канала
	var mu sync.Mutex
	go func() {
		for val := range result {
			mu.Lock()
			sum += val
			mu.Unlock()
		}
	}()

	// горутины пишут значения в канал и добавляются в WaitGroup
	for i := 0; i < count; i++ {
		workers <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-workers
			}()
			result <- 1
		}()
	}

	wg.Wait()
	// mutex, waitingGroup, буферизированные каналы и все равно без слипа нет 100% гарантии что будет выведено верно
	time.Sleep(200 * time.Millisecond)
	lgr.Logger(fmt.Sprintf("result: %d", sum))
}
