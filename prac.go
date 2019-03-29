package main

import "fmt"

func main() {
	b := true
	opt := ""
	switch b {
	case 1 < 2:
		fmt.Println("ashchi")
		opt = "yoyo"
	case 6 < 8:
		fmt.Println("yeay")
		opt = "nyo nyo"

	}
	fmt.Println(opt)

}
