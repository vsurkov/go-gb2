package main

import (
	"os"
	"sync"
)

type File struct {
	dir   string
	finfo os.FileInfo
	dinfo os.DirEntry
}

type Result struct {
	files struct {
		sync.RWMutex
		m map[string]File
	}
	//files  map[string]File
	dupl struct {
		sync.RWMutex
		m map[string]map[string]File
	}

	// fcount, dcount int64
	// tfs, tds int64 // total files size, total duplicates size
}

func NewStorage() *Result {
	var s Result
	s.files.m = make(map[string]File)
	s.dupl.m = make(map[string]map[string]File)
	return &s
}
