package main

import (
	"log"
	"time"
)

// Supply function for calculating runtime of func()
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
