package main

import (
	"fmt"
	"github.com/namsral/flag"
	"hash/crc32"
	"io"
	"log"
	"os"
	"strings"
)

//
//Программа поиска дубликатов файлов.
//Дубликаты файлов - это файлы, которые совпадают по имени файла и по его размеру.

//1.Программа должна работать на локальном компьютере и получать на вход путь до директории.
//2.Программа должна вывести в стандартный поток вывода список дублирующихся файлов:
//2.1 которые находятся как в директории
//2.2 поддиректориях директории, переданной через аргумент командной строки.
//3. Программа должна работать эффективно при помощи распараллеливания программы.
//4. Программа должна принимать дополнительный ключ - возможность удаления обнаруженных дубликатов файлов после поиска.
//5. Дополнительнно нужно обезопасить пользователей от случайного удаления файлов.

func main() {
	var path string
	var rm bool

	// Инициализируем переменные
	flag.StringVar(&path, "path", ".", "set folder path to scan")
	flag.BoolVar(&rm, "rm", false, "remove duplicate files")

	log.Println(path)
	log.Println(rm)

	//Запускем обработку файловой структуры
	Crawl(path)
}

func Crawl(dir string) {
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for i := range dirEntry {
		log.Println(entryFormatter(dir, dirEntry[i]))
		if dirEntry[i].IsDir() {
			Crawl(strings.Join([]string{dir, dirEntry[i].Name()}, "/"))
		}
	}
}

func entryFormatter(dir string, entry os.DirEntry) string {
	var hash string

	if !entry.IsDir() {
		hash = calcHash(dir + "/" + entry.Name())
	}
	return fmt.Sprintf("%s/%s\t%s", dir, entry.Name(), hash)
}

func calcHash(fullName string) string {
	const chunkSize = 10

	//Открываем файл, обрабатываем ошибки и ошибки закрытия и ошибки записи ошибки в os.Stderr
	f, err := os.Open(fullName)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "can't close file %v", err.Error())
		if err != nil {
			log.Println("Error on write error on os.Stderr")
		}
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			_, err = fmt.Fprintf(os.Stderr, "can't close file %v", err.Error())
			if err != nil {
				log.Println("Error on write error on os.Stderr")
			}
		}
	}(f)

	//Читаем файл чанками в байтовый массив
	buf := make([]byte, chunkSize)

	for {
		_, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Println(err)
		}

		if err == io.EOF {
			break
		}
	}

	h := crc32.NewIEEE()
	_, err = h.Write(buf)

	return fmt.Sprintf("%v", h.Sum32())
}
