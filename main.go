package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered panic:", v)
		}
	}()
	err := router()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
