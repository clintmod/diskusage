package main

import (
	//"fmt"
	"os"
	"log"
	"sync"
	"fmt"
	"io/ioutil"
)

func main() {
	dirList(os.Args[1])
}

func dirList(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var float_size float64
	var size_unit string
	var file_name string
	var file_path string
	for _, file := range files {
		file_name = file.Name()
		file_path = path + file_name
		if (!file.IsDir()) {
			float_size, size_unit = getFileSize(file.Size())
		} else {
			var size, _= dirSize(file_path)
			float_size, size_unit = getFileSize(size)
		}
		fmt.Printf("%.2f%s\t%s\n", float_size, size_unit, file_path)
	}
}

func dirSize(path string) (int64, error) {
	var mu sync.Mutex
	var size int64 = 0
	err := fastWalk(path, func(path string, typ os.FileMode) error {
		mu.Lock()
		defer mu.Unlock()
		if !typ.IsDir() {
			var f, err = os.Lstat(path)
			if err != nil {
				return err
			}
			size += f.Size()
		}
		return nil
	})
	return size, err
}

func getFileSize(size int64)(float64, string){
	var float_size float64 = float64(size)
	var size_unit string = "B"
	if(float_size > 1024) {
			float_size = float_size / 1024
			size_unit = "K"
		}
	if(float_size > 1024) {
			float_size = float_size / 1024
			size_unit = "M"
		}
	if(float_size > 1024) {
			float_size = float_size / 1024
			size_unit = "G"
		}
	return float_size, size_unit
}