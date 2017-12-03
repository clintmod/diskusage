package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	find(os.Args[1])
}

func find(path string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	// Get file names
	files, err := f.Readdirnames(0)
	if err != nil {
		return
	}
	sort.Strings(files)
	// show files
	for _, v := range files {
		fmt.Println(v)
	}
}

