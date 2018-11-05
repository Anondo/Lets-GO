package main

import (
	"fmt"
	"strconv"
)

func main() {
	myInt := 12
	myFloat := 3.5
	myStr1 := "100"
	myStr2 := "50.8"
	/* int -> float  */
	describe(float64(myInt))
	/* float -> int  */
	describe(int(myFloat))
	/* string -> int */
	strInt, _ := strconv.ParseInt(myStr1, 0, 0)
	describe(strInt)
	/* string -> float */
	strFloat, _ := strconv.ParseFloat(myStr2, 64)
	describe(strFloat)
	/* int -> string */
	intStr := strconv.Itoa(myInt)
	describe(intStr)
	/* float -> string */
	floatStr := strconv.FormatFloat(myFloat, 'f', 2, 64)
	describe(floatStr)

}

func describe(a interface{}) {
	fmt.Printf("Type of %v is %T\n", a, a)
}
