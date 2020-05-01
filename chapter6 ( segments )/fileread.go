package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("chapter6 ( segments )/test.txt")
	var strings []string

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tmp := scanner.Text()

		if tmp == "" {
			continue
		}

		strings = append(strings, tmp)
	}

	fmt.Println(strings[5:8])
}
