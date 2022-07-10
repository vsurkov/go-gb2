package main

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func Scan(dir string) *Result {
	return filter(walk(filepath.Clean(dir)))
}

// Упрощенный и ускоренный поиск дубликатов, на основе размера файла и имени
func walk(dir string) *Result {
	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	workers := make(chan struct{}, runtime.NumCPU()*2)

	wg.Add(1)
	go func(dir string) {
		workers <- struct{}{}
		defer func() { <-workers }()
		mux.Lock()
		dirEntry, err := os.ReadDir(dir)
		mux.Unlock()
		errorHandler("can't read dir", err)

		wg2 := sync.WaitGroup{}
		subWorkers := make(chan struct{}, runtime.NumCPU()*2)
		for i := range dirEntry {
			wg2.Add(1)
			go func(i int) {
				subWorkers <- struct{}{}
				defer func() { <-subWorkers }()
				if dirEntry[i].IsDir() {
					walk(filepath.Join(dir, dirEntry[i].Name()))
				}
				entryProcess(dir, dirEntry[i])
				wg2.Done()
			}(i)
		}
		wg2.Wait()
		wg.Done()
	}(dir)
	wg.Wait()
	return search
}

// Из предподготовленного списка дубликатов нужно отфильтровать случайные совпадения на основе CRC-суммы
func filter(search *Result) *Result {
	unfiltered := search.dupl.m
	var filtered = make(map[string]map[string]File)

	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	workers := make(chan struct{}, runtime.NumCPU()*8)

	// Разберем мапу нефильтрованых с ключом по хешу
	for hash, m := range unfiltered {
		wg.Add(1)
		go func(hash string, m map[string]File) {
			workers <- struct{}{}
			defer func() { <-workers }()
			var crcMaster string
			// Разберем вложенную мапу файлов с ключом по текущей директории
			mux.Lock()
			for dir, file := range m {
				if crcMaster == "" {
					crcMaster = getMD5hash(filepath.Join(file.dir, file.finfo.Name()))
				}
				crcFile := getMD5hash(filepath.Join(file.dir, file.finfo.Name()))

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
			mux.Unlock()
			wg.Done()
		}(hash, m)
	}
	wg.Wait()
	close(workers)
	// заполним очищенный результат фильтрованым
	search.dupl.m = removeLonely(filtered)

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
