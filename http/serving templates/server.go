package main

import (
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

func makePlayer(name string, club string, age int) Player {
	return Player{name, club, age}
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)

	players := make([]Player, 3)
	players = append(players, makePlayer("Messi", "Fc Barcelona", 31))
	players = append(players, makePlayer("C.Ronaldo", "Real Madrid Cf", 34))
	players = append(players, makePlayer("Neyman", "PSG", 27))

	var context struct {
		Footballers []Player
	}
	context.Footballers = players

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, context)
}
func dummyhandler(w http.ResponseWriter, r *http.Request) {}
