package main

import "fmt"

type Node struct{
    next *Node
    data int
}
func nodeConstructor(d int , n *Node) *Node{
    newNode := new(Node)
    newNode.next = n
    newNode.data = d
    return newNode
}

type List struct{
    head *Node
}
func listConstructor() *List{
    newlist := new(List)
    newlist.head = nodeConstructor(0 , nil)
    return newlist
}

type LinkedList interface{
    push(int)
    showList()
    pop() int
}
func (l *List) push(data int){
    temp := nodeConstructor(data , l.head)
    l.head = temp
}
func (l List) showList(){
    var current = l.head
    for current.next != nil{
        fmt.Println(current.data)
        current = current.next
    }
}
func (l *List) pop() int{
    data := l.head.data
    l.head = l.head.next
    return data
}

func main(){
    var ll LinkedList = listConstructor()
    ll.push(12)
    ll.push(11)
    ll.push(10)
    ll.push(9)
    ll.showList()
    fmt.Println("Popped Value",ll.pop())
    ll.showList()
}
