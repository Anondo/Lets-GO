package main

import "fmt"

func main(){
    characters := make(map[string]string) //specifying the index type and the type of the values this container will hold
    episodes := make(map[int]string) //make function is used to make slices or maps

    characters["king in the north"] = "Jon Snow"
    characters["lady of winterfell"] = "Sansa Stark"
    characters["mother of dragons"] = "Daenerys Targaryan"

    episodes[1] = "Episode 1"
    episodes[2] = "Episode 2"
    episodes[4] = "Episode 4"

    for rank := range characters{
        fmt.Println(characters[rank] ,"is" , rank)
    }
    for episode , name := range episodes{
        fmt.Println("Episode ", episode ," is " , name)
    }

    charName , exists := characters["king in the north"]
    if(exists){
        fmt.Println(charName , " is  a fine character")
    }else{
        fmt.Println("Who is this")
    }

    delete(characters , "king in the north") //delete key

    charName , exists = characters["king in the north"]
    if(exists){
        fmt.Println(charName , " is  a fine character")
    }else{
        fmt.Println("Who is this")
    }

    characters["lady of winterfell"] = "Arya Stark" //update

    for rank := range characters{
        fmt.Println(characters[rank] ,"is" , rank)
    }

}

//maps are sort of like dictionary in python but they are not dynamic. They can only contain one type of value
