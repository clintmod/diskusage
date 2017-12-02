package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	dirList(os.Args[1])
}

func dirList(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if(!file.IsDir()) {
			fmt.Printf("name: %s, size: %.2f KB\n", file.Name(), float64(file.Size()) / 1024)
		} else {
			var size, file_count = dirSize(path + "/" + file.Name())
			var float_size float64 = float64(size) / 1024
			var size_unit string = "KB"
			if(float_size > 1024) {
				float_size = float_size / 1024
				size_unit = "MB"
			}
			if(float_size > 1024) {
				float_size = float_size / 1024
				size_unit = "GB"
			}
			fmt.Printf("name: %s, size: %.2f %s, count:%v\n", file.Name(), float_size, size_unit, file_count)
		}
	}
}

func dirSize(path string)(int64, int64) {
	var size int64
	var file_count int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
			file_count++
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return size, file_count
}
