package main

import (
	"fmt"
	"time"
)

func gorutinePanic() {
	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recovered", v)
			}
		}()
		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}
