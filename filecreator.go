package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func createBillionFiles(path string) error {
	err := validatePath(path)
	if err != nil {
		fmt.Println("error: %s", err)
		return err
	}

	err = validateDir(path)
	if err != nil {
		return err
	}

	return nil
}

func validatePath(dir string) error {
	if valid := fs.ValidPath(dir); valid != true {
		return errors.New("dir not valid")
	}
	return nil
}

func validateDir(dir string) error {
	root := dir
	fileSystem := os.DirFS(root)
	//fmt.Println(getCurrentPath())

	err := fs.WalkDir(fileSystem, ".", func(dir string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	fmt.Printf("Path '%s' is valid\n", dir)
	return nil
}

func getCurrentPath() string {
	dir, err := os.Getwd()
	check(err)
	return dir
}

func createFile(path string, prefix string) error {
	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
