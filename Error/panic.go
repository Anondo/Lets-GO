/* recover is a function used to handle error caused by a panic */
/* when a panic is raised, the flow of control goes to the defer from the normal functions */
package main

import "fmt"

func dividePos(arg1, arg2 int) int {
	defer func() {
		fmt.Println(recover()) //printing the error message
	}() //call to defer must be a function

	if arg1 < 0 || arg2 < 0 {
		panic("Got negative argument(s)") //casuing a panic on purpose
	}

	return arg1 / arg2

}

func main() {
	fmt.Println(dividePos(5, 0))
	fmt.Println(dividePos(6, -2))
	fmt.Println(dividePos(6, 2))
}
