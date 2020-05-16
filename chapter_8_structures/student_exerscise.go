package main

import "fmt"

type student struct {
	name  string
	grade float64
}

func printInfo(s *student) {
	s.grade += 1.0
	fmt.Println("Name \t:", s.name)
	fmt.Printf("Grade\t: %0.1f\n", s.grade)
}

func main() {
	var s student
	s.name = "ALonzo Cole"
	s.grade = 92.3
	printInfo(&s)
}
