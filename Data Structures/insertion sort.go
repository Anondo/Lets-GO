package main

import "fmt"

func insertionSort(arr []int) []int{
    for i := 0 ; i < len(arr) -1; i++{
        if(arr[i] > arr[i+1]){
            arr = swap(i , i+1 , arr)
            arr = subCheck(i , arr)
        }
    }
    return arr
}
func swap(fIndex int , lIndex int , arr []int) []int{
    temp := arr[fIndex]
    arr[fIndex] = arr[lIndex]
    arr[lIndex] = temp
    return arr
}
func subCheck(index int , arr []int) []int{
    for i := index-1; i >= 0;  i--{
        if(arr[i] > arr[index]){
            arr = swap(i , index , arr)
            index = i
        }
    }
    return arr
}

func main(){
    myArr := []int{33,11,1,57,24,3}
    myArr = insertionSort(myArr)
    fmt.Println(myArr)
}
