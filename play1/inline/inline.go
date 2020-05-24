package main

import (
	"fmt"
	"os"
)

func main() {
	if n, err := fmt.Printf("Hello, World!\n"); err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("The number of bytes returned: %d\n", n)
	}
}
