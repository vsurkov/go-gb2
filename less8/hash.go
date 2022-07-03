package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func calcHash(fullName string) string {
	const chunkSize = 10

	//Открываем файл, обрабатываем ошибки и ошибки закрытия и ошибки записи ошибки в os.Stderr
	f, err := os.Open(fullName)
	errorHandler("can't close file", err)

	defer func(f *os.File) {
		err := f.Close()
		errorHandler("can't close file", err)
	}(f)

	//Читаем файл чанками в байтовый массив
	buf := make([]byte, chunkSize)

	for {
		_, err := f.Read(buf)
		if err != nil && err != io.EOF {
			errorHandler("error reading chunk of file", err)
		}

		if err == io.EOF {
			break
		}
	}

	h := md5.New()
	_, err = h.Write(buf)
	errorHandler("error on write on hash", err)

	return fmt.Sprintf("%x", h.Sum(nil))
}
