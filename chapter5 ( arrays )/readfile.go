package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, error := os.Open("/home/yevgen/go/src/test/chapter5 ( arrays )/data.txt")

	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("", scanner.Text())
	}

	error = file.Close()

	if error != nil {
		log.Fatal(error)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}

func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64

	file, error := os.Open(fileName)

	if error != nil {
		return numbers, error
	}

}
