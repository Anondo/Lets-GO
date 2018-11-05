package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("myfile.txt")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some real aweosome text")
	file.Close()

	stream, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type of stream is %T\n", stream)

	text := string(stream)

	fmt.Println(text)

}
