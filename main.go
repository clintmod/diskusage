package main

import (
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
	var floatSize float64
	var sizeUnit string
	var fileName string
	var filePath string
	for _, file := range files {
		fileName = file.Name()
		filePath = path + fileName
		if (!file.IsDir()) {
			floatSize, sizeUnit = getFileSize(file.Size())
		} else {
			var size, _= dirSize(filePath)
			floatSize, sizeUnit = getFileSize(size)
		}
		fmt.Printf("%.1f%s\t%s\n", floatSize, sizeUnit, filePath)
	}
}

func dirSize(path string) (int64, error) {
	var mu sync.Mutex
	var size int64
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
	var floatSize = float64(size)
	var sizeUnit = "B"
	if(floatSize > 1024) {
			floatSize = floatSize / 1024
			sizeUnit = "K"
		}
	if(floatSize > 1024) {
			floatSize = floatSize / 1024
			sizeUnit = "M"
		}
	if(floatSize > 1024) {
			floatSize = floatSize / 1024
			sizeUnit = "G"
		}
	return floatSize, sizeUnit
}