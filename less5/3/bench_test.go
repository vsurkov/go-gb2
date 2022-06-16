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
	fmt.Printf("Mutex set:\n")
	shaperMux(10, 1000, b)
}
func BenchmarkRWSetUsecaseOne(b *testing.B) {
	fmt.Printf("RWMutex set:\n")
	shaperRWMux(10, 1000, b)
}

////50% запись, 50% чтение;
//func BenchmarkUsecaseTwo(b *testing.B) {
//	shaperMux(50, 1000, b)
//	shaperRWMux(50, 1000, b)
//}
//
////90% запись, 10% чтение
//func BenchmarkUsecaseThree(b *testing.B) {
//	shaperMux(90, 1000, b)
//	shaperRWMux(90, 1000, b)
//}

func shaperMux(writePercent int, scope int, b *testing.B) {

	writeScope := scope * (writePercent / 100)
	readScope := scope * (1 - (writePercent / 100))

	SetAdd(b, writeScope)
	SetHas(b, readScope)
}

func shaperRWMux(writePercent int, scope int, b *testing.B) {

	writeScope := scope * (writePercent / 100)
	readScope := scope * (1 - (writePercent / 100))

	RWSetAdd(b, writeScope)
	RWSetHas(b, readScope)
}

func SetAdd(b *testing.B, parallels int) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(parallels)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

func SetHas(b *testing.B, parallels int) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(parallels)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func RWSetAdd(b *testing.B, parallels int) {
	var set = NewRWSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(parallels)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

func RWSetHas(b *testing.B, parallels int) {
	var set = NewRWSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(parallels)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}
