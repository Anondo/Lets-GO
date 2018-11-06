package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	fmt.Fprintf(w, "Hello world....first golang server...yeay")
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
