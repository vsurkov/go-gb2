package main

import (
	"fmt"
	"log"
	"os"
)

func errorHandler(text string, err error) {
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "%v: %v\n", text, err.Error())
		if err != nil {
			log.Println(err.Error())
		}
	}
}
