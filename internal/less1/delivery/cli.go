package less1

import (
	"fmt"
	less1 "github.com/vsurkov/go-gb2/internal/less1/internal/less1/usecase"
)

func Router() error {
	var choice int

	msg := `ДЗ 1. Введите целое число в диапазоне от 1 до 4, включительно:
	1 - функция с неявной паникой
	2 - функция с неявной паникой и возвратом кастомной ошибки
	3 - создание миллиона пустых файлов
	4 - задание с паникой в горутине 
Ваш выбор: `
	fmt.Println(msg)

	_, err := fmt.Scanln(&choice)
	if err != nil {
		return err
	}

	switch choice {
	case 1:
		less1.SimplePanic()
		return nil
	case 2:
		if err := less1.CustomPanicError(); err != nil {
			return err
		}
		return nil
	case 3:
		err = less1.CreateFiles(1000)
		if err != nil {
			return err
		}
		return nil
	case 4:
		less1.GoroutinePanic()
		return nil

	default:
		return nil
	}
}
