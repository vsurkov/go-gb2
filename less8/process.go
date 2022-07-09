package main

import (
	"os"
)

var search = NewStorage()

func entryProcess(dir string, entry os.DirEntry) {
	var hash string

	if entry.IsDir() {
		return
	}

	hash = getSizeNameHash(entry)

	// Проверить если в хранилище уже есть искомый хеш то добавить найденный ранее файл и новый в список дублей
	oldFile, exist := search.files[hash]
	entryInfo, _ := entry.Info()
	if exist { // если вложенная в dupl нет значения по hash - нужно инициализировать вложенную мапу
		dummy, ok := search.dupl[hash]
		if !ok {
			dummy = make(map[string]File)
			search.dupl[hash] = dummy
		}

		// сохраняем в дубликаты
		search.dupl[hash][oldFile.dir] = oldFile
		search.dupl[hash][dir] = File{
			dir:   dir,
			finfo: entryInfo,
			dinfo: entry,
		}
	}

	//
	search.files[hash] = File{
		dir:   dir,
		finfo: entryInfo,
		dinfo: entry,
	}

}
