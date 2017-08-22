package main

import ("fmt"
        "errors")


func main(){
    answer , err := divide(2,0)
    if(err != nil){
        fmt.Println(err)
    }else{
        fmt.Println(answer)
    }
}


func divide(arg1 float64, arg2 float64) (float64 , error){
    if(arg2 == 0){
        return 0 , errors.New("Something divided by zero is undefined looser!!")
    }
    return arg1/arg2 , nil
}
