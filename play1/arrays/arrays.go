package main

import "fmt"

func printer(words [4]string) {
	for _, word := range words {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")

	// make modification to the array
	words[2] = "blue"
}

func main() {

	// array of strings of unknown length
	words := [...]string{"the", "quick", "brown", "fox"}
	printer(words)
	printer(words)
}
