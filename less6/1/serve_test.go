package main

import (
	"sync"
	"testing"
)

func TestRunOnce(t *testing.T) {
	runServe(t, 100, 100, 4950)
}

func TestRunOneListener(t *testing.T) {
	runServe(t, 1000, 1, 499500)
}

func TestRunOneCount(t *testing.T) {
	runServe(t, 1, 100, 0)
}

func TestRepeat(t *testing.T) {
	runRepeatServe(t, 10000, 100, 1000, 49995000000)
}

func runServe(t *testing.T, count int, listeners int, expected int) {
	received := serve(count, listeners)
	if received != expected {
		t.Errorf("Run with count: %d, listeners: %d was failed: expected result %d but result is: %d",
			count,
			listeners,
			expected,
			received)
	}
}

func runRepeatServe(t *testing.T, count int, listeners int, repeats int, expected int) {
	received := 0

	var mux = sync.Mutex{}
	for i := 0; i < repeats; i++ {
		tmp := serve(count, listeners)
		mux.Lock()
		received += tmp
		mux.Unlock()
	}

	if received != expected {
		t.Errorf("Run with count: %d, listeners: %d was failed: expected result %d but result is: %d",
			count,
			listeners,
			expected,
			received)
	}
}
