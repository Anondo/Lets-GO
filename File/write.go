package main

import (
	"fmt"
	"log"
	"os"
)

func appendToFile(filename, text string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(text)
}

func main() {
	appendToFile("file.txt", "This is the first line\n")
	appendToFile("file.txt", "This is the second line\n")
	fmt.Println("Done writing")
}
