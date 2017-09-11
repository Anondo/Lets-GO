package main

import "fmt"

func main(){
    var countryMap = map[string]string{"Bangladesh" : "Dhaka" , "England" : "London" , "Spain" : "Madrid"} //map literal (no need to use make as no nil value)
    for country , capital := range countryMap{
        fmt.Println("The capital of",country,"is",capital)
    }
}
