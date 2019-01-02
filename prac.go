package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Object map[string]string

func main() {
	s := "{\"en\":\"50% Discount on next 5 rides!!!\"}"
	b := []byte(s)
	var o Object
	err := json.Unmarshal(b, &o)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)

}
