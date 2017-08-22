package main

import "fmt"

func main(){
    players := []string {"Messi","Neymar","Suarez","Ronaldo","Dyabla"}

    for index , value := range players{
        fmt.Println(value," is the no.",index+1," player")
    }
}

//range returns the index and value as a pair in an expression
