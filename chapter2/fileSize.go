package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, readError := os.Stat("/home/yevgen/go/src/test/chapter2/fileSize.go")
	if readError != nil {
		log.Fatal(readError)
	}

	fmt.Println(fileInfo.Name(), fileInfo.Size())
}
