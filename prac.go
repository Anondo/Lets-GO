package main

import (
	"fmt"
	"log"
	"net/http"

	randu "github.com/thedevsaddam/renderer"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	rndr := randu.New()
	rndr.HTMLString(rw, http.StatusOK, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not start server: %v", err))
	}

}
