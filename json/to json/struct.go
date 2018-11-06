package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id     int      `json:"id,omitempty"` //tags determines how json output , omitempty ensures empty data ar omitted
	Name   string   `json:"name,omitempty"`
	Emails []string `json:"emails,omitempty"`
}

func main() {
	allen := User{1, "Barry Allen", []string{"barry@dc.com", "theflash@jlu.com"}}
	fmt.Println(getJSON(allen))
}

func getJSON(u User) string {
	json_bytes, _ := json.Marshal(u)
	return string(json_bytes)
}
