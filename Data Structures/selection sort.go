package main

import "fmt"

func selection_sort(arr []int) []int{
    var minIndex int
    for i:=0; i < len(arr)-1; i++{
        minIndex = findMinIndex(i , arr[i+1:])
        if(arr[i] > arr[minIndex]){
            arr = swap(i , minIndex , arr)
        }
    }
    return arr
}
func findMinIndex(curIndex int , arr []int) int{
    minIndex := 0
    for i:=0; i < len(arr); i++{
        if(arr[i] < arr[minIndex]){
            minIndex = i
        }
    }
    return minIndex + curIndex + 1
}
func swap(index1 int , index2 int , arr []int) []int{
    temp := arr[index1]
    arr[index1] = arr[index2]
    arr[index2] = temp
    return arr
}

func main(){
    myArr := []int{14,33,27,10,35,19,42,44}
    myArr = selection_sort(myArr)
    fmt.Println(myArr)
}
