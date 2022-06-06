package main

import (
	"fmt"
	less1 "github.com/vsurkov/go-gb2/internal/less1/internal/less1/delivery"
	"os"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered panic:", v)
		}
	}()
	err := less1.Router()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
