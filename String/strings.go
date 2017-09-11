package main

import ("fmt"
        "strings")

func main(){
    str :=  "Hello World"
    fmt.Println(str)
    fmt.Println(len(str))
    for i:= 0; i < len(str); i++{
        fmt.Print(string(str[i]))
    }
    fmt.Println()
    fmt.Println(strings.Split(str , " "))
    fmt.Println(strings.Split(str , " ")[0])

}
