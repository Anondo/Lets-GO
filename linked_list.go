package main

import "fmt"

func main(){
    var list linked_list
    list.push(12)
    list.push(11)
    fmt.Println("Hello World")

}

type Stack interface{
    push(int)
    //pop()
}

type linked_list struct{
    head node
}
type node struct{
    next linked_list
    data int
}

func (l linked_list) push(data int){
    var temp node
    temp.data = data
    temp.next = l.head
    l.head = temp
}
