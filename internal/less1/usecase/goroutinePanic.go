/*
 Учебный пример обработки паник в горутине
*/
package less1

import (
	"fmt"
	"time"
)

//GoroutinePanic() реализует простой пример обработки паники в горутине
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
