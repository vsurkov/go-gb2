package main

import (
	"fmt"
	"testing"
)

//Протестируйте производительность операций чтения и записи на множестве
//действительных чисел, безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов
//использования:

//10% запись, 90% чтение;
func BenchmarkSetUsecaseOne(b *testing.B) {
	muxSet(b, 10, 1000)
}

func BenchmarkRWSetUsecaseOne(b *testing.B) {
	rwMuxSet(b, 10, 1000)
}

////50% запись, 50% чтение;
func BenchmarkSetUsecaseTwo(b *testing.B) {
	muxSet(b, 50, 1000)
}
func BenchmarkRWSetUsecaseTwo(b *testing.B) {
	rwMuxSet(b, 50, 1000)
}

////90% запись, 10% чтение
func BenchmarkSetUsecaseThree(b *testing.B) {
	muxSet(b, 90, 1000)
}
func BenchmarkRWSetUsecaseThree(b *testing.B) {
	rwMuxSet(b, 90, 1000)
}

func muxSet(b *testing.B, writePercent float32, scope float32) {
	var set = NewSet()
	var s = NewScope(writePercent, scope)
	writeScope := s.GetWriteScope()
	readScope := s.GetReadScope()
	fmt.Printf("\nMutex set, writeScope:%v, readScope:%v of total:%v\n",
		writeScope,
		readScope,
		scope)

	b.ResetTimer()
	b.StartTimer()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(writeScope)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})

	b.Run("", func(b *testing.B) {
		b.SetParallelism(readScope)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
	b.StopTimer()
}

func rwMuxSet(b *testing.B, writePercent float32, scope float32) {
	var set = NewRWSet()
	var s = NewScope(writePercent, scope)
	writeScope := s.GetWriteScope()
	readScope := s.GetReadScope()
	fmt.Printf("\nRWMutex set, writeScope:%v, readScope:%v of total:%v\n",
		writeScope,
		readScope,
		scope)

	b.ResetTimer()
	b.StartTimer()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(writeScope)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})

	b.Run("", func(b *testing.B) {
		b.SetParallelism(readScope)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
	b.StopTimer()
}
