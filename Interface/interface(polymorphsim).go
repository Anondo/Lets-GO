package main

import "fmt"

type Shape interface{
    getArea() float64
}

type Circle struct{
    radius float64
}
type Triangle struct{
    base , height float64
}
type Rectangle struct{
    length , width float64
}

func (c Circle) getArea() float64{
    return (3.1416) * c.radius * c.radius
}
func (t Triangle) getArea() float64{
    return (0.5) * t.base * t.height
}

func main(){
    var shape Shape
    c := Circle{2.8}
    t := Triangle{4.7,3}
    //r := Rectangle{6,4.9}
    shape = c
    fmt.Println("Area of cirle",shape.getArea())
    shape = t
    fmt.Println("Area of Triangle" , shape.getArea())
    //shape = r
    //fmt.Println("Area of Rectangle" , shape.getArea())
}

//uncomment the above lines and we will get an error because rectangle doesnt implement Shape by using all its methods
