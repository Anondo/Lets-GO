package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFromFile(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func main() {

	fmt.Println(readFromFile("file.txt"))
}
