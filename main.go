package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"sort"
	"crypto/md5"
	"io"
	"runtime"
	"github.com/stretchr/powerwalk"
)

func main() {
	var cpus = 2//runtime.NumCPU()
	fmt.Println(cpus)
	runtime.GOMAXPROCS(cpus)
	var path = os.Args[1]
	dirList(path)
	//dirList2(path)
	/*if b, err := ComputeMd5(path); err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("%x", b)
	}*/
}

func ComputeMd5(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	return hash.Sum(result), nil
}

func dirList2(path string) {
	var file, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	var names, _ = file.Readdirnames(0)
	sort.Strings(names)
	for _, name := range names {
		var info, error = os.Lstat(path + "/" + name)
		if error != nil {
			log.Fatal(error)
		}
		var float_size, size_unit = getFileSize(info.Size())
		fmt.Printf("name: %s, size: %.2f %v\n", name, float_size, size_unit)
	}
}

func dirList(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if(!file.IsDir()) {
			var float_size, size_unit = getFileSize(file.Size())
			fmt.Printf("name: %s, size: %.2f %v\n", file.Name(), float_size, size_unit)
		} else {
			var size, _ = dirSize(path + "/" + file.Name())
			var float_size, size_unit = getFileSize(size)
			fmt.Printf("name: %s, size: %.2f %s\n", file.Name(), float_size, size_unit)
		}
	}
}

func dirSize(path string)(int64, int64) {
	var size int64
	var file_count int64
	err := powerwalk.Walk(path, func(_ string, info os.FileInfo, err error) error {
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

func getFileSize(size int64)(float64, string){
	var float_size float64 = float64(size)
	var size_unit string = "B"
	if(float_size > 1024) {
		float_size = float_size / 1024
		size_unit = "KB"
	}
	if(float_size > 1024) {
		float_size = float_size / 1024
		size_unit = "MB"
	}
	if(float_size > 1024) {
		float_size = float_size / 1024
		size_unit = "GB"
	}
	return float_size, size_unit
}