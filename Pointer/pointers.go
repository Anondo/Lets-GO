package main

import "fmt"

func main(){
     a := 12
     var p *int //pointer declaration
     p = &a

     fmt.Println("Address of a = " , &a)
     fmt.Println("Value of a = " , a)
     fmt.Println("Value of p = " , p)
     fmt.Println("Value of address p is pointing = " , *p)
      fmt.Println("Address of p = " , &p)

     var pp **int //pointer to a pointer
     pp = &p

     fmt.Println("Value of pp(address of the pointer it is pointing to) = " , pp)
     fmt.Println("Value of *pp(value of the pointer it is pointing to) = " , *pp)
     fmt.Println("Value of **pp(value of the variable the pointer is pointing which it is pointing to!!) = " , **pp)

     x := 12
     y := 10
     fmt.Println("Before swap:  X = " , x , " Y = " , y)
     swap(&x , &y)
     fmt.Println("After swap:  X = " , x , " Y = " , y)
}

func swap(a *int , b *int){ //does require to return the changed values as they are changed by reference 
    temp :=  *a
    *a = *b
    *b = temp
}
