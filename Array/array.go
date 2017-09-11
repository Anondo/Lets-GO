package main

import "fmt"

func main(){
    var arr1 [5]int
    var arr2 = [5]float64{1.1 , 2.2 , 3.3 , 4.4 , 5.5}
    var slice = []int {2 , 4 , 6 , 8 , 10} //this is a slice not an array..slice looks array but it is more like python array(dynamic)
    var slice2 = []int {}

    for i:=0; i < len(arr1); i++{
        arr1[i] = i * i
    }
    fmt.Println(arr1)
    fmt.Println(arr2)
    for i:=0; i < len(arr1); i++{
        fmt.Print(arr1[i] , "\t")
    }
    fmt.Println()
    half(slice)
    fmt.Println(slice)
    half2(arr1)
    fmt.Println(arr1)
    slice2 = arrInit(slice)
    fmt.Println(slice2)
}

func half(arr []int){ //only slice can be passed through a func like this....the reference is passed
    for i:=0; i < len(arr); i++{
        arr[i] = arr[i]/2
    }
}
func half2(arr [5]int){
    for i:=0; i < len(arr); i++{
        arr[i] = arr[i]/2
    }
}
func arrInit(arr []int) []int{
    for i:=0; i < len(arr); i++{
        arr[i] = i * i * i
    }
    return arr
}
