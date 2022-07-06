package main

import (
	"os"
	"strings"
)

func Scan(dir string) *Storage {
	dirEntry, err := os.ReadDir(dir)
	errorHandler("can't read dir", err)

	for i := range dirEntry {
		if dirEntry[i].IsDir() {
			Scan(strings.Join([]string{dir, dirEntry[i].Name()}, "/"))
		}
		entryProcess(dir, dirEntry[i])
	}
	return search
}
