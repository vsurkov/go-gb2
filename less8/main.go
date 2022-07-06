package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/namsral/flag"
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
	flag.StringVar(&path, "p", ".", "set path to scan")
	flag.BoolVar(&rm, "rm", false, "remove duplicate files")
	flag.Parse()

	//Запускем сканирование файловой структуры
	//res := Scan(path)
	res := Scan("/Users/HOMEr/Downloads")

	//var totalSize
	for key, val := range res.dupl {
		var size uint64
		fcount := uint64(len(val))

		fmt.Printf(DebugColor, fmt.Sprintf("HASH %v\n", key))
		for kk, vv := range val {
			fmt.Printf("  %v/%v\n", kk, vv.finfo.Name())

			if size == 0 {
				size = uint64(vv.finfo.Size())
			}

		}

		fmt.Printf(InfoColor,
			fmt.Sprintf("Files: %v\tSize: %v \tTotal: %v",
				fcount,
				humanize.Bytes(size),
				humanize.Bytes(size*fcount)))
		fmt.Printf(ErrorColor,
			fmt.Sprintf("\t Overspending: %v\n\n", humanize.Bytes(size*(fcount-1))))
	}
}

//entryInfo, err := entry.Info()
//errorHandler("getting oldFile Info error", err)
//size := uint64(entryInfo.Size())
//name := entryInfo.Name()
