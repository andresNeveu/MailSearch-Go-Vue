package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// take first command line argument, path
	pathArg := os.Args[1]

	// get directory list
	innerPath := "maildir"
	dirPath := filepath.Join(pathArg, innerPath)
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		check(err)
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)

		return nil
	})
	check(err)
}
