package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := map[string]interface{}{
		"id":     1,
		"name":   "Barry Allen",
		"emails": []string{"barry@dc.com", "allen@jlu.com"},
	}
	fmt.Println(getJSON(user))
}

func getJSON(data map[string]interface{}) string {
	json_bytes, _ := json.Marshal(data)
	return string(json_bytes)
}
