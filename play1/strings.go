package main

import "fmt"

func main() {

	atoz := "the quick brown fox jumps over the lazy dog\n"

	// take substrings
	fmt.Printf("%s\n", atoz[15:])

	// Going through a string rune by rune
	// We do that by using a for loop
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	// Get length of the string
	fmt.Printf("Length of the string: %d\n", len(atoz))
}
