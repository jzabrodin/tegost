package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	number := rand.Intn(1000)
	attempts := 15

	lowBorder := 1
	highBorder := 10000

	getUserInput(int64(number), attempts, lowBorder, highBorder)
}

func getUserInput(number int64, attempts int, lowBorder int, highBorder int) {

	if attempts == 0 {
		fmt.Println("Game over! Number is : ", number)
		os.Exit(1)
	}

	fmt.Println("Attempts count:", attempts)
	fmt.Println("Please enter a number from ", lowBorder, " to ", highBorder)
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	data = strings.TrimSpace(data)

	if err != nil {
		log.Fatal(err)
	}

	intData, err := strconv.ParseInt(data, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	if intData == number {
		fmt.Println("You win!")
	} else if intData > number {
		fmt.Println("Your guess is HIGH")
		getUserInput(number, attempts-1, lowBorder, int(intData))
	} else if intData < number {
		fmt.Println("Your guess is LOW")
		getUserInput(number, attempts-1, int(intData), highBorder)
	}
}
