package main

import "fmt"

type Author struct {
	Name string
	Bio  string
}

func (a *Author) GetBio() string {
	bio := a.Bio
	return bio
}

type Book struct {
	Author   //composition happens here
	BookName string
	BookType string
}

func (b *Book) GetBio() string { // kinda like method overriding
	bio := b.Author.GetBio() //kinda like super
	return bio
}

func main() {
	book := Book{Author{"Allen", "Awesome writer"}, "The greatest book", "Knowledge"}
	fmt.Println(book.GetBio())
}
