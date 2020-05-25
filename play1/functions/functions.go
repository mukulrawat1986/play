package main

import (
	"fmt"
	"os"
)

func filePrinter(msg string) (err error) {

	f, err := os.Create("helloworld.txt")
	if err != nil {
		return
	}
	defer f.Close()

	f.Write([]byte(msg))
	return
}

func printer(format string, msgs ...string) {
	defer fmt.Printf("\n")
	for _, msg := range msgs {
		fmt.Printf(format, msg)
	}
}

func main() {

	printer("%s", "Hello, ", " World!", " This", " is", " dog") // print string
	printer("%x", "Hello, ", " World!", " This", " is", " dog") // print in hex

}
