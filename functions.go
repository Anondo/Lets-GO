package main

import "fmt"

func main(){
    a := 8
    b := 10
    a = square(a)
    b = square(b)
    fmt.Println(a)
    fmt.Println(b)
    a , b = swap(a , b)
    fmt.Println(a)
    fmt.Println(b)
}

func square(num int) int{
    return num*num
}

func swap(arg1 , arg2 int) (int , int){
    return arg2 , arg1
}
