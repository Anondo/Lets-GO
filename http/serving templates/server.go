package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Player struct {
	Name string
	Club string
	Age  int
}

func main() {
	http.HandleFunc("/favicon.ico", dummyhandler)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)

	var players []Player
	fmt.Println(players)
	players = append(players, Player{"Messi", "Fc Barcelona", 31})
	players = append(players, Player{"C.Ronaldo", "Real Madrid Cf", 34})
	players = append(players, Player{"Neyman", "PSG", 27})

	var context struct {
		Footballers []Player
	}
	context.Footballers = players

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, context)
}
func dummyhandler(w http.ResponseWriter, r *http.Request) { log.Println(r.Method) }
