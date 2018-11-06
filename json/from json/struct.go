package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id     int      `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Emails []string `json:"emails,omitempty"`
}

func main() {

	flash := Person{}
	setData(&flash)
	fmt.Println(flash)

}

func setData(p *Person) {
	data := []byte(`
          {
            "Id":2,
            "Name":"Barry Allen",
            "Emails":["barry@dc.com","theflash@jlu.com"]
          }
    `)
	err := json.Unmarshal(data, &p)
	defer func() {
		fmt.Println(recover())
	}()
	if err != nil {
		panic(err)
	}

}
