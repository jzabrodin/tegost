package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	names := getNames()
	ranks := make(map[string]int)
	var namesUnique []string
	for _, name := range names {
		if ranks[name] == 0 {
			ranks[name] = 1
			namesUnique = append(namesUnique, name)
		}

		ranks[name] = ranks[name] + 1
	}

	fmt.Println("before", namesUnique)
	sort.Strings(namesUnique)
	fmt.Println("after:", namesUnique)

	for _, val := range namesUnique {
		fmt.Println(val, ":", ranks[val])
	}
}

func getNames() []string {
	file, err := os.Open("chapter 7 ( maps )/votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	var names []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	return names
}
