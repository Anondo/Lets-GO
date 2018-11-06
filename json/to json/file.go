package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Emails []string `json:emails`
}

func main() {
	flash := Person{2, "Barry Allen", []string{"barry@dc.com", "theflash@jlu.com"}}
	json_file, err := os.OpenFile("output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(json_file).Encode(flash)
	fmt.Println("Done")
}
