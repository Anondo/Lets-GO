package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	id, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[1], 0, 0)
	if err != nil {

	}
	fmt.Fprintf(w, "Your id is "+strconv.Itoa(int(id)))
}

func main() {
	port := 8080
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server running at http://localhost:%v\n", port)
}
