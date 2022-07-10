package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func getMD5hash(fullName string) string {
	const chunkSize = 1_024_000

	// Открываем файл, обрабатываем ошибки и ошибки закрытия и ошибки записи ошибки в os.Stderr
	f, err := os.Open(fullName)
	errorHandler("can't close file", err)

	// Читаем файл чанками в байтовый массив
	buf := make([]byte, chunkSize)

	for {
		_, err := f.Read(buf)
		if err != nil {
			break
		}
	}

	h := md5.New()
	_, err = h.Write(buf)
	errorHandler("error on hash write", err)
	defer func(f *os.File) {
		err := f.Close()
		errorHandler("can't close file", err)
	}(f)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getSizeNameHash(entry os.DirEntry) string {
	entryInfo, err := entry.Info()
	errorHandler("getting file Info error", err)

	s := fmt.Sprintf("%v %v", entryInfo.Size(), entryInfo.Name())
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
