package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	person := make(map[string]interface{})
	setData(person)
	fmt.Println(person)

}

func setData(person map[string]interface{}) {
	data := []byte(`
          {
            "Id":2,
            "Name":"Barry Allen",
            "Emails":["barry@dc.com","theflash@jlu.com"]
          }
    `)
	err := json.Unmarshal(data, &person)
	defer func() {
		fmt.Println(recover())
	}()
	if err != nil {
		panic(err)
	}
}
