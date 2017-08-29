package main

import "fmt"

type Node struct{
    leftChild *Node
    rightChild *Node
    data int
}
type tree struct{
    root *Node
}
func BFTConstructor() *tree{
    temp := new(tree)
    temp.root = new(Node)
    return temp
}
type BinarySearchTree interface{
    insert(int)
    search(int)
    postTraverse(*Node)
    postOrderTraversal()
    preTraverse(*Node)
    preOrderTraversal()
}
func (t *tree) insert(data int){
    if(t.root.data == 0){
        t.root.data = data
    }else{
        parent := t.root
        current := parent
        for true{
            if(parent.leftChild == nil || parent.rightChild == nil){
                parent.leftChild = new(Node)
                parent.rightChild = new(Node)
            }
            if(data < current.data){
                current = current.leftChild
            }else{
                current = current.rightChild
            }
            if(current.data == 0){
                current.data = data
                break
            }else{
                parent = current
            }
        }
    }
}
func (t *tree) search(data int){
    current := t.root
    if(current.data == 0){
        fmt.Println("Tree is Empty")
        return
    }
    for current.data != data{
        fmt.Println("Comparing with", current.data)
        if(data < current.data){
            fmt.Println("Moving Left")
            current = current.leftChild
        }else{
            fmt.Println("Moving Right")
            current = current.rightChild
        }
        if(current == nil){
            fmt.Println("Data Not Found")
            return
        }
    }
    fmt.Println("Found")
}
func (t *tree) postTraverse(current *Node){
        if(current.data != 0){
            if(current.leftChild != nil || current.rightChild != nil){
                t.postTraverse(current.leftChild)
                t.postTraverse(current.rightChild)
                fmt.Println(current.leftChild.data)
                fmt.Println(current.rightChild.data)
            }
            if(current == t.root){
                fmt.Println(current.data)
            }
        }
}
func (t *tree) postOrderTraversal(){
    t.postTraverse(t.root)
}
func (t *tree) preTraverse(current *Node){
        if(current.data != 0){
            fmt.Println(current.data)
            if(current.leftChild != nil || current.rightChild != nil){
                t.preTraverse(current.leftChild)
                t.preTraverse(current.rightChild)
            }
        }
}
func (t *tree) preOrderTraversal(){
    t.preTraverse(t.root)
}


func main(){
    var bft BinarySearchTree = BFTConstructor()
    bft.insert(27)
    bft.insert(14)
    bft.insert(35)
    bft.insert(10)
    bft.insert(19)
    bft.insert(31)
    bft.insert(42)
    bft.search(20)
    bft.postOrderTraversal()
    bft.preOrderTraversal()
}
