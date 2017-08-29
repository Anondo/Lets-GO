package main

import "fmt"

func main(){
    arr := []int{11,4,17,5,1,3}
    fmt.Println(arr)
    arr = bubble_sort(arr)
    fmt.Println(arr)
}


func bubble_sort(arr []int) []int{
    for i:=0; i < len(arr); i++{
        for j:=0; j < len(arr) - i - 1; j++{
            if(arr[j] > arr[j+1]){
                temp := arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = temp
            }
        }
    }
    return arr
}
