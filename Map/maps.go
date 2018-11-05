package main

import "fmt"

func main() {
	characters := make(map[string]string) //specifying the index type and the type of the values this container will hold
	episodes := make(map[int]string)      //make function is used to make slices or maps

	characters["king in the north"] = "Jon Snow"
	characters["lady of winterfell"] = "Sansa Stark"
	characters["mother of dragons"] = "Daenerys Targaryan"

	episodes[1] = "Winter is coming"
	episodes[2] = "Something something"
	episodes[4] = "The red wedding"

	for rank, character := range characters {
		fmt.Println(character, "is the ", rank)
	}
	for episode, name := range episodes {
		fmt.Println("Episode ", episode, " is ", name)
	}

	charName, exists := characters["king in the north"]
	if exists {
		fmt.Println(charName, " is  a fine character")
	} else {
		fmt.Println("Who is this")
	}

	delete(characters, "king in the north") //delete key

	charName, exists = characters["king in the north"]
	if exists {
		fmt.Println(charName, " is  a fine character")
	} else {
		fmt.Println("Who is this")
	}

	characters["lady of winterfell"] = "Arya Stark" //update

	for rank := range characters {
		fmt.Println(characters[rank], "is", rank)
	}

}

//maps are sort of like dictionary in python but they are not dynamic. They can only contain one type of value
