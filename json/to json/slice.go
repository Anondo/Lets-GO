package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	email := []string{"barry@dc.com", "theflash@jlu.com"}
	fmt.Println(getJSON(email))
}

func getJSON(email []string) string {
	json_bytes, _ := json.Marshal(email)
	return string(json_bytes)
}
