package main

import (
	"fmt"
)

func router() error {
	var choice int

	msg := `ДЗ 1. Введите целое число в диапазоне от 1 до 4, включительно:
	1 - функция с неявной паникой
	2 - функция с неявной паникой и возвратом кастомной ошибки
	3 - создание миллиона пустых файлов
	4 - задание с паникой в горутине
Ваш выбор: `
	fmt.Printf(msg)

	_, err := fmt.Scanln(&choice)
	if err != nil {
		return err
	}

	switch choice {
	case 1:
		simplePanic()
		return nil
	case 2:
		if err := customPanicError(); err != nil {
			return err
		}
		return nil
	case 3:
		var path string
		fmt.Print("Укажите имя директории для создания файлов: ")
		_, err := fmt.Scanln(&path)
		if err != nil {
			return err
		}

		err = validatePath(path)
		if err != nil {
			fmt.Println("error: %s", err)
			return err
		}

		return nil

	default:
		return nil
	}
}
