package main

import "os"

type File struct {
	dir   string
	finfo os.FileInfo
	dinfo os.DirEntry
}

type Storage struct {
	files map[string]File
	dupl  map[string]map[string]File
	//fcount, dcount int64
	//tfs, tds int64 // total files size, total duplicates size
}

func NewStorage() *Storage {
	var s Storage
	s.files = make(map[string]File)
	s.dupl = make(map[string]map[string]File)
	return &s
}
