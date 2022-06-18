package main

import (
	"fmt"
	"sync"
)

//3. Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”

func main() {

	var wgl sync.WaitGroup
	var wgw sync.WaitGroup

	var count = 1000
	var listeners = 10

	var wp = make(chan struct{}, listeners)
	var data = make(chan int, 3)

	var sum int

	for l := 0; l < listeners; l++ {
		wgl.Add(1)
		go func() {
			defer wgl.Done()
			for value := range data {
				//Гонка здесь
				sum += value
				//Гонка закончилась
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

	fmt.Println(sum)
}
