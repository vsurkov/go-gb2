package less1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	path = "/tmp/gb/dat"
)

//Простая реализация, которая падает.
//Recovered panic: open /tmp/gb/dat253: too many open files
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

// CreateFiles рабочий и самый быстрый из трех стандартных способов записи данных при создании файла
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

//Результат запуска - ошибки нет:
//На 1000000 2022/06/06 16:09:38 create files took 3m35.382211087s
//На 10000 2022/06/06 16:03:37 create files took 1.185854488s

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
//2022/06/06 00:20:05 create files took 2h4m48.764971194s
//ls /tmp/gb | wc -l
//1000001

//На 10000 штуках 2022/06/06 15:09:06 create files took 1m14.357327911s

func createStringFile(n int) (wb int, err error) {
	fullPath := path + strconv.Itoa(n)
	file, err := os.Create(fullPath)
	check(err)
	defer func() {
		err = file.Close()
		check(err)
	}()

	//Пишем стринг
	//wb, err = file.WriteString("I will not burp in class.\n")
	wb, err = file.WriteString("ooooo")
	check(err)
	err = file.Sync()
	check(err)
	return wb, nil
}

//На 10000 штуках 2022/06/06 15:49:34 create files took 1m13.40286839s
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s заняло %s", name, elapsed)
}
