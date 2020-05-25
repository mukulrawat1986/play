package main

import "os"

func printer(msg string) (err error) {

	f, err := os.Create("helloworld.txt")
	if err != nil {
		return
	}
	defer f.Close()

	f.Write([]byte(msg))
	return
}

func main() {

	printer("Hello, World!")

}
