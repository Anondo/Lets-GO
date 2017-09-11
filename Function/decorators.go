package main

import "fmt"

func main(){
    square := decorate(square)
    fmt.Println(square(8))

}

func square(num int) int{
    return num*num
}

func decorate(f func(int) int) func(int) int{
    return func(num int) int{
        return f(num) + 1
    }
}
