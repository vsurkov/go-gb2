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
	search.files.Lock()
	oldFile, exist := search.files.m[hash]
	search.files.Unlock()
	entryInfo, _ := entry.Info()
	if exist { // если вложенная в dupl нет значения по hash - нужно инициализировать вложенную мапу
		search.dupl.Lock()
		dummy, ok := search.dupl.m[hash]
		search.dupl.Unlock()
		if !ok {
			dummy = make(map[string]File)
			search.dupl.Lock()
			search.dupl.m[hash] = dummy
			search.dupl.Unlock()
		}

		// сохраняем в дубликаты
		search.dupl.Lock()
		search.dupl.m[hash][oldFile.dir] = oldFile
		search.dupl.m[hash][dir] = File{
			dir:   dir,
			finfo: entryInfo,
			dinfo: entry,
		}
		search.dupl.Unlock()
	}

	search.files.Lock()
	search.files.m[hash] = File{
		dir:   dir,
		finfo: entryInfo,
		dinfo: entry,
	}
	search.files.Unlock()
}
