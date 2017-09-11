package main

import ("fmt"
    "bufio"
    "os")


func main(){
    reader := bufio.NewReader(os.Stdin)
    var name string
    name  , _  = reader.ReadString('\n')
    fmt.Println("Hello" , name)
}
