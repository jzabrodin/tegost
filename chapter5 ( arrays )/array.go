package main

import "fmt"

func main() {
	var notes [7]string

	notes[0] = "Do"
	notes[1] = "Re"
	notes[2] = "Mi"
	notes[3] = "Fa"
	notes[4] = "So"
	notes[5] = "La"
	notes[6] = "Si"

	fmt.Println("Notes: ", notes)

	var notes2 [7]string = [7]string{"d", "r", "m", "f", "s", "l", "s"}

	fmt.Println("", notes2, "length", len(notes2), notes2[6])

	for index, value := range notes2 {
		fmt.Println("", index, " ", value)
	}

}
