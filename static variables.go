package main
import "fmt"

func main(){
    var a = 12 //during compile time the  compiler decides the type of the variable depending on the type of the value
    var b int = 13 //declareing the type after the variable name is optional
    var f1 , f2  float64
    f1 = 1.2
    f2 = 2.5
    var s string
    s = "Hello"
    var s2 = " World"
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(f1)
    fmt.Println(f2)
    fmt.Println(s + s2)
}
