package main

import "fmt"
import "time"

func getInt(a chan interface{}) {
	time.Sleep(100 * time.Microsecond)
	a <- 12
}
func getStr(a chan interface{}) {
	//time.Sleep(100 * time.Microsecond)
	a <- "Hello world"
}
func describe(i interface{}) {
	fmt.Printf("%v is a %T\n", i, i)
}

func main() {
	ch := make(chan interface{})
	go getInt(ch)
	go getStr(ch)
	x, y := <-ch, <-ch
	describe(x)
	describe(y)
}
