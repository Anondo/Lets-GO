package main

import "fmt"

type Books struct{
     id string
     name string
     author string
     copies int
}

func (b Books) getName() string{
    return b.name
}


func main(){
    var book1 Books
    book1.id = "11A09B6"
    book1.name = "Harry Potter"
    book1.author = "J.K. Rowling"
    book1.copies = 100
    book2 := Books {"22HJK89" , "Sherlock Holmes" , "Sir Arthur Conan Doyle" , 1000}
    fmt.Println(book1.id)
    fmt.Println(book1.name)
    fmt.Println(book1.author)
    fmt.Println(book1.copies)
    fmt.Println(book2.id)
    fmt.Println(book2.name)
    fmt.Println(book2.author)
    fmt.Println(book2.copies)
    fmt.Println(book1.getName())
    fmt.Println(book2.getName())
}
