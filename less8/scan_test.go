package main

import (
	"fmt"
	"log"
	"testing"
)

func TestScanLists(t *testing.T) {
	res := Scan("/Users/HOMEr/Downloads/test2")

	//for key, val := range res.dupl {
	//	fmt.Printf("HASH %v\n", key)
	//	for kk, vv := range val {
	//		fmt.Printf("%v/%v\n", kk, vv.finfo.Name())
	//	}
	//	fmt.Println()
	//}
	exp := 11
	rec := len(res.files.m)
	if exp != rec {
		t.Errorf("Expected %v, but received %v", exp, rec)
	}

	exp = 1
	rec = len(res.dupl.m)
	if exp != rec {
		t.Errorf("Expected %v, but received %v", exp, rec)
	}
	fmt.Println()
}

func TestScanResult(t *testing.T) {
	result := Scan("/Users/HOMEr/Downloads/test2")

	files := make(map[string]string)
	files["ce8b86f8ab8248fe4025b725a974b410"] = ".DS_Store"
	files["e8653eb7d5b21873e59e21327f3c39c6"] = "2.pdf"
	files["c5ad1c2c80372864b96eaef66df91553"] = "3.pdf"
	files["e3ca8db27329bf2a7cda0034b39daf2a"] = "4 — копия.pdf"
	files["34d5f3819ec89f5d54bc9dfc70ea16ba"] = "7.pdf"
	files["8022917250c52fab4c931d740ecbd049"] = "4.pdf"
	files["74ec4d375e9c854a0768e2a81a975167"] = "foobar.txt"
	files["88ca03c702968c7e9b3623a023fca3c5"] = "0.jpeg"
	files["78c57f5a6aca63527737dd0131c84af0"] = "1.zip"
	files["e8d3c69aff7d453b462c799f00330181"] = "5.pdf"
	files["ce603004a31f43d6385a649943423d2f"] = "6.pdf"

	for key := range files {
		exp := files[key]

		mapVal, ok := result.files.m[key]
		if !ok {
			t.Errorf("result.files[%v] empty", key)
			log.Fatal()
		}
		rec := mapVal.finfo.Name()

		if exp != rec {
			t.Errorf("Expected %v, but received %v", exp, rec)
		}
	}

	dupl := make(map[string]string)
	duplKey := "88ca03c702968c7e9b3623a023fca3c5"
	dupl["/Users/HOMEr/Downloads/test2"] = "0.jpeg"
	dupl["/Users/HOMEr/Downloads/test2/folder"] = "0.jpeg"

	for key := range dupl {
		exp := dupl[key]

		mapVal, ok := result.dupl.m[duplKey][key]
		if !ok {
			t.Errorf("result.dupl[%v][%v] empty", duplKey, key)
			log.Fatal()
		}

		rec := mapVal.finfo.Name()
		if exp != rec {
			t.Errorf("Expected %v, but received %v", exp, rec)
		}
	}
}
