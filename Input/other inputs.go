package main

import "fmt"

func main() {
    var arg1 , arg2 float64
    var opt string
    fmt.Print("Enter first argument: ")
    fmt.Scanf("%f\n" , &arg1)
    fmt.Print("Enter second argument: ")
    fmt.Scanf("%f\n" , &arg2)
    fmt.Print("Enter  operation: ")
    fmt.Scanf("%s\n" , &opt)
    if(opt == "+"){
        fmt.Printf("%f + %f = %f" , arg1 , arg2 , arg1+arg2)
    }else if(opt == "-")
        fmt.Printf("%f - %f = %f" , arg1 , arg2 , arg1-arg2)
    }else if(opt == "*"){
        fmt.Printf("%f * %f = %f" , arg1 , arg2 , arg1*arg2)
    }else if(opt == "/"){
        fmt.Printf("%f / %f = %f" , arg1 , arg2 , arg1/arg2)
    }else{
        fmt.Println("Invalid Operation")
    }
}
