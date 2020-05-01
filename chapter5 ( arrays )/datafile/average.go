package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filepath := "/Users/yevgen-zabrodin/Documents/tegost/chapter5 ( arrays )/data.txt"
	values, err := GetFloats(filepath)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	fmt.Println(values)

	for _, number := range values {
		sum += number
	}
	fmt.Println("average : ", sum/float64(len(values)))
}

// GetFloats : getting floats from a file
func GetFloats(filename string) ([4]float64, error) {

	var numbers [4]float64
	file, err := os.Open(filename)

	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)

	i := 0

	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return numbers, err
		}
		i++
	}

	err = file.Close()

	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil

}
