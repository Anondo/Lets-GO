package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var hero map[string]interface{}
	json_file, err := os.Open("input.json")
	if err != nil {
		panic(err)
	}
	json.NewDecoder(json_file).Decode(&hero)
	fmt.Println(hero)
}
