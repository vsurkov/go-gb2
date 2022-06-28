package main

import (
	"log"
)

//Написать функцию, которая принимает на вход имя файла и название функции.
//Необходимо подсчитать в этой функции количество вызовов асинхронных функций.
//Результат работы должен возвращать количество вызовов int и ошибку error.
//Разрешается использовать только go/parser, go/ast и go/token.

func main() {
	file := "parseFile.go"
	goroutines, err := parseFile(file, "parseFile")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(goroutines)
}
