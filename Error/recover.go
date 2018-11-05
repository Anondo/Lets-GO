/* recover is a function used to handle error caused by a panic */
/* when a panic is raised, the flow of control goes to the defer from the normal functions */
package main

import "fmt"

func divide(arg1, arg2 int) int {
	defer func() {
		fmt.Println(recover()) //printing the error message
	}() //call to defer must be a function

	return arg1 / arg2

}

func main() {
	fmt.Println(divide(5, 0))
	fmt.Println(divide(6, 2))
}
