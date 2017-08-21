package main

import "fmt"

func main(){
    nterm := 10
    fmt.Println(nterm,"th term: " , fibonacci(nterm))
    fmt.Println("Values upto " , nterm , "th term: ")
    for i := 0 ; i <= nterm; i++{
        fmt.Println(fibonacci(i))
    }
}

func fibonacci(n int) int{
    if(n == 0){
        return 0
    }
    if(n == 1){
        return 1
    }
    return fibonacci(n - 1) + fibonacci(n-2)
}
