/*
 Реализация механизма создания заданного количества файлов в файловой системе
*/

package less1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

//Константа содержит в себе дефаултный путь к директории внутри которого будут созданы файлы
const (
	path = "/tmp/gb/dat"
)

//Простая реализация, которая падает при запуске. Она не может быть использована в реальной жизни - пример.
//Recovered panic: open /tmp/gb/dat253: too many open files
//функция принимает аргумент count int равное количеству создаваемых файлов
//dummyCreateFiles(count int) error
//возвращает error
func dummyCreateFiles(count int) error {
	fileCount := count
	defer timeTrack(time.Now(), "create files")
	for n := 0; n <= fileCount; n++ {
		fullPath := path + strconv.Itoa(n)
		_, err := os.Create(fullPath)
		check(err)
	}
	return nil
}

//CreateFiles в реализации использует createBuffFile() буферизированный вывод как самый быстрый при создании
//функция принимает аргумент **count int** равный количеству создаваемых файлов
//CreateFiles(count int) error
//возвращает error в случае ошибки, в журнал пишет информацию о продолжительности выполнения
//и количестве созданых файлов
func CreateFiles(count int) error {
	var totalBytes int
	fileCount := count
	defer timeTrack(time.Now(), "INFO: создание файлов")
	for n := 0; n <= fileCount; n++ {
		wb, err := createBuffFile(n)
		if err != nil {
			return err
		}
		totalBytes += wb

	}
	log.Printf("INFO: в каталоге %s успешно создано %d файлов \n", path, count)
	return nil
}

//createBuffFile рабочий и самый быстрый из трех стандартных способов записи данных при создании файла,
//функция принимает аргумент **n int** порядковый номер для имени, в реализации использует буферизированный вывод
//createBuffFile(n int) (wb int, err error)
// возвращает количество записанных байт и error
func createBuffFile(n int) (wb int, err error) {
	fullPath := path + strconv.Itoa(n)
	file, err := os.Create(fullPath)
	check(err)
	defer func() {
		err = file.Close()
		check(err)
	}()

	//Пишем используя буфер
	w := bufio.NewWriter(file)
	wb, err = w.WriteString("ooooo")
	check(err)
	err = w.Flush()
	check(err)

	return wb, nil
}

//Результат запуска - ошибки нет:
//На 1000000 2022/06/06 16:09:38 create files took 3m35.382211087s
//На 10000 2022/06/06 16:03:37 create files took 1.185854488s

//createStringFile отличается от предыдущего варианта отсутствием буфера при записи, что драматически снижает быстродействие
//функция принимает аргумент **n int** порядковый номер для имени, в реализации использует буферизированный вывод
//createStringFile(n int) (wb int, err error)
// возвращает количество записанных байт и error
func createStringFile(n int) (wb int, err error) {
	fullPath := path + strconv.Itoa(n)
	file, err := os.Create(fullPath)
	check(err)
	defer func() {
		err = file.Close()
		check(err)
	}()

	//Пишем стринг
	wb, err = file.WriteString("I will not burp in class.\n")
	check(err)
	err = file.Sync()
	check(err)
	return wb, nil
}

//Результат запуска - ошибки нет:
//На 1000000 2022/06/06 00:20:05 create files took 2h4m48.764971194s
//На 10000 2022/06/06 15:09:06 create files took 1m14.357327911s

//createBytesFile отличается от createStringFile способом записи в файл используя массив байт не влияет на скорость
//функция принимает аргумент **n int** порядковый номер для имени, в реализации использует буферизированный вывод
//createBytesFile(n int) (wb int, err error) {
// возвращает количество записанных байт и error

func createBytesFile(n int) (wb int, err error) {
	fullPath := path + strconv.Itoa(n)
	file, err := os.Create(fullPath)
	check(err)
	defer func() {
		err = file.Close()
		check(err)
	}()

	//Пишем байты
	data := []byte{111, 111, 111, 111, 111}
	wb, err = file.Write(data)
	check(err)
	err = file.Sync()
	check(err)

	return wb, nil
}

//Результат запуска - ошибки нет:
//На 10000 штуках 2022/06/06 15:49:34 create files took 1m13.40286839s

//Функция check(e error) используется сокращения кода, реализует под капотом проверку на наличие ошибки
//и генерацию паники - Panic() используется только для учебных целей
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Функция timeTrack() используется для подсчета времени исполнения метода
//считается как дельта между
//defer timeTrack(time.Now(), "INFO: создание файлов")
//и временем выхода из функции
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s заняло %s", name, elapsed)
}
