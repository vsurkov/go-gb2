package main

import (
	"runtime"
	"testing"
)

var count = 10000
var listeners = runtime.NumCPU()

func BenchmarkServe1(b *testing.B) {
	serve(count, 1)
}
func BenchmarkServe2(b *testing.B) {
	serve(count, 1)
}

func BenchmarkServeX05(b *testing.B) {
	serve(count, listeners/2)
}
func BenchmarkServeX(b *testing.B) {
	serve(count, listeners)
}

func BenchmarkServeX2(b *testing.B) {
	serve(count, listeners*2)
}

func BenchmarkServeX4(b *testing.B) {
	serve(count, listeners*4)
}

func BenchmarkServeX8(b *testing.B) {
	serve(count, listeners*8)
}

func BenchmarkServeX16(b *testing.B) {
	serve(count, listeners*16)
}

func BenchmarkServe1k(b *testing.B) {
	serve(count, 1000)
}

func BenchmarkServe10k(b *testing.B) {
	serve(count, 10000)
}
