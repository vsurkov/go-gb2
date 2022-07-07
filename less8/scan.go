package main

import (
	"fmt"
	"os"
	"strings"
)

func Scan(dir string) *Result {
	return filter(scan(dir))
}

// Упрощенный и ускоренный поиск дубликатов, на основе размера файла и имени
func scan(dir string) *Result {
	dirEntry, err := os.ReadDir(dir)
	errorHandler("can't read dir", err)

	for i := range dirEntry {
		if dirEntry[i].IsDir() {
			scan(strings.Join([]string{dir, dirEntry[i].Name()}, "/"))
		}
		entryProcess(dir, dirEntry[i])
	}
	return search
}

// Из предподготовленного списка дубликатов нужно отфильтровать случайные совпадения на основе CRC-суммы
func filter(search *Result) *Result {
	unfiltered := search.dupl
	var filtered = make(map[string]map[string]File)

	// Разберем мапу нефильтрованых с ключом по хешу
	for hash, m := range unfiltered {
		var crcMaster string
		// Разберем вложенную мапу файлов с ключом по текущей директории
		for dir, file := range m {
			if crcMaster == "" {
				crcMaster = getMD5hash(fmt.Sprintf("%v/%v", file.dir, file.finfo.Name()))
			}
			crcFile := getMD5hash(fmt.Sprintf("%v/%v", file.dir, file.finfo.Name()))

			// Если crc текущего файла совпадает с мастер значит это дубликат, иначе отбрасываем
			if crcMaster == crcFile {
				// если нет значения по hash - нужно инициализировать вложенную мапу
				dummy, ok := filtered[hash]
				if !ok {
					dummy = make(map[string]File)
					filtered[hash] = dummy
				}

				// сохраняем в список фильтрованого мапы
				filtered[hash][dir] = file
			}
			delete(m, dir)
		}
		delete(unfiltered, hash)
	}
	// заполним очищенный результат фильтрованым
	search.dupl = removeLonely(filtered)

	return search
}

// Нужно очистить от мап с одним файлом
func removeLonely(search map[string]map[string]File) map[string]map[string]File {
	for hash := range search {

		if len(search[hash]) == 1 {
			delete(search, hash)
		}
	}
	return search
}
