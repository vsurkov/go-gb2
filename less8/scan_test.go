package main

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	res := Scan("/Users/HOMEr/Downloads/test2")

	for key, val := range res.dupl {
		fmt.Printf("HASH %v\n", key)
		for kk, vv := range val {
			fmt.Printf("%v/%v\n", kk, vv.finfo.Name())
		}
		fmt.Println()
	}

	fmt.Println(len(res.files))
	fmt.Println(len(res.dupl))
}
