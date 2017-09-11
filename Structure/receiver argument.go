package main

import "fmt"

type Coordinate struct{
    x , y float64
}

func (c Coordinate) getX() float64{ //this is a receiver argument which is a special type of argument for a function(making it a method) which takes the struct type for which it will be used
    return c.x
}
func (c Coordinate) getY() float64{
    return c.y
}
func main(){
    c := Coordinate{2.5 , 7.8}
    fmt.Println(c.getX())
    fmt.Println(c.getY())
}
