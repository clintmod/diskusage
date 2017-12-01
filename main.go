package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gosuri/uilive"
	"github.com/stretchr/powerwalk"
)

const kb = 1024.00

var writer = uilive.New()
var size float64
var filecount uint64

func main() {
	defer timeTrack(time.Now(), "main")
	writer.Start()
	dirSize(os.Args[1])
	writer.Stop()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func dirSize(path string) (float64, error) {
	err := powerwalk.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += float64(info.Size())
			filecount++
			log(size, filecount)
		}
		return err
	})
	return size, err
}

var lastLogTime = time.Now()

func log(size float64, filecount uint64) {
	currentTime := time.Now()
	if currentTime.Sub(lastLogTime) < time.Millisecond*110 {
		return
	}
	fmt.Fprintf(writer, "Size: %f GB | File count: %d\n", size/(kb*kb*kb), filecount)
	writer.Flush()
	lastLogTime = currentTime
}
