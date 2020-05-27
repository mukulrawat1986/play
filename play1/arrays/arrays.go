package main

import "fmt"

func printer(words []string) {
	for _, word := range words {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")

}

func main() {

	// slice of strings of unknown length
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	fmt.Printf("Length of string: %d\n", len(words))

	printer(words[:2])
	printer(words[2:])
	printer(words[2:4])
}
