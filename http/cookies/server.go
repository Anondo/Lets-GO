package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func loginHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie, _ := r.Cookie("username")
		if cookie != nil {
			fmt.Fprintf(rw, "Hi %v,  you are already logged in", cookie.Value)
		} else {
			t, _ := template.ParseFiles("login.html")
			t.Execute(rw, struct{}{})
		}
	}
	if r.Method == "POST" {
		r.ParseForm()
		username := strings.ToLower(r.Form["username"][0])
		password := r.Form["password"][0]
		if username == "bob" && password == "supersecret" {
			expiration := time.Now().Add(24 * time.Minute)
			cookie := http.Cookie{Name: "username", Value: username, Expires: expiration}
			http.SetCookie(rw, &cookie)
			http.Redirect(rw, r, "/home", http.StatusOK)

		} else {
			fmt.Fprintln(rw, "Invalid credentials")
		}
	}
}

func logoutHandler(rw http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("username")
	cookie := http.Cookie{Name: c.Name, Value: "", Expires: time.Unix(0, 0)}
	http.SetCookie(rw, &cookie)
	http.Redirect(rw, r, "/login", http.StatusOK)
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	if cookie == nil {
		fmt.Fprintln(rw, "You have to log in to watch this page")
	} else {
		context := struct{ Username string }{cookie.Value}
		t, _ := template.ParseFiles("homepage.html")
		t.Execute(rw, context)
	}

}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/logout", logoutHandler)
	log.Println("Starting server at localhost:8080... ... ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
