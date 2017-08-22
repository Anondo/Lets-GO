package main

import "fmt"

type Shape interface{
    area() float64
}

type Circle struct{
    radius float64
}
type Triangle struct{
    length , width float64
}

func (t Triangle) area() float64{ //implemenetaion of shape interface by triangle struct
    return (0.5) * t.length * t.width
}
func (c Circle) area() float64{
    return (3.1416) * c.radius * c.radius
}
func  getArea(s Shape) float64{ //both circle and triangle is now a shape as they have implemented all the methods of the shape interface
    return s.area()
}

func main(){
    circle := Circle{2.5}
    triangle := Triangle{1.6,5.8}
    fmt.Println("Area of the circle: " ,  getArea(circle))
    fmt.Println("Area of the triangle: " ,  getArea(triangle))
}

//in go interface are implemneted by structures implicitly. There is no "implements key-word". The only way to implement is to implement all
//the methods
