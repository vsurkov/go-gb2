package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

//1. Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
//Выполните трассировку программы

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var count = 10000
	var listeners = runtime.NumCPU() * 8

	fmt.Print(serve(count, listeners))
}
