package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("current time", time.Now().YearDay())
	broken := "G# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println("fixed:", fixed)
}
