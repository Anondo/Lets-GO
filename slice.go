package main

import "fmt"

func main(){
    arr := []int{} //this is a slice not an array(no element count)
    //var arr []int ...this is the same as above
    arr = append(arr , 1)
    arr = append(arr , 2)
    arr = append(arr , 3)
    arr = append(arr , 4)
    fmt.Println(arr)
    fmt.Println(arr[1:3])
}
 //slices are abstraction of array. They are like dynamic arrays, similar to that of python
