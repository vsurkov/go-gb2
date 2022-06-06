package less1

import (
	"fmt"
	"time"
)

func GoroutinePanic() {
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
