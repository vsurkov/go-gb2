package main

import (
	"fmt"
	humanize "github.com/dustin/go-humanize"
	"log"
	"os"
	"strings"
)

func Scan(dir string) {
	dirEntry, err := os.ReadDir(dir)
	errorHandler("can't read dir", err)

	for i := range dirEntry {
		log.Println(entryFormatter(dir, dirEntry[i]))
		if dirEntry[i].IsDir() {
			Scan(strings.Join([]string{dir, dirEntry[i].Name()}, "/"))
		}
	}
}

func entryFormatter(dir string, entry os.DirEntry) string {
	var hash string

	if entry.IsDir() {
		return fmt.Sprintf("%s %s/%s", hash, dir, entry.Name())
	}

	hash = calcHash(dir + "/" + entry.Name())

	entryInfo, err := entry.Info()
	errorHandler("getting file Info error", err)
	size := uint64(entryInfo.Size())
	name := entryInfo.Name()

	return fmt.Sprintf("%s\t%v\t%s/%s", hash, humanize.Bytes(size), dir, name)
}
