package main

import (
	"fmt"
	"log"
)

// Написать функцию, которая принимает на вход структуру
// in (struct или кастомную struct) и
// values map[string]interface{}, где
// (key - название поля структуры, которому нужно присвоить value этой мапы).
// Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
// Функция может возвращать только ошибку error.
// Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Weight float64 `json:"weight"`
	Male   bool    `json:"male"`
	Place  struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
}

func (p Person) print() {
	fmt.Printf("name: %v\n", p.Name)
	fmt.Printf("age: %v\n", p.Age)
	fmt.Printf("weight: %v\n", p.Weight)
	fmt.Printf("male: %v\n", p.Male)
	//fmt.Printf("place:\n")
	//fmt.Printf("\tlatitude: %v\n", p.Place.Latitude)
	//fmt.Printf("\tlongitude: %v\n", p.Place.Longitude)
}

func main() {
	bob := Person{
		Name:   "bob",
		Age:    20,
		Weight: 88.1,
		Male:   true,
		Place: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}(struct {
			Latitude  float64
			Longitude float64
		}{Latitude: 55.751244, Longitude: 37.618423}),
	}

	changes := map[string]interface{}{
		"Name":   "jane",
		"Age":    22,
		"Weight": 90.9,
		"Male":   false,
	}

	err := peopleChanger(&bob, changes)
	if err != nil {
		log.Fatal(err)
	}
}
