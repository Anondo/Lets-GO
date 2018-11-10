package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, req *http.Request) {
   if req.Method == "GET"{
     tmplt  , err := template.ParseFiles("templates/login.html")
     if err != nil{
       log.Fatal(err)
     }
     tmplt.Execute(rw , nil)
   }else if req.Method == "POST"{
     tmplt  , err := template.ParseFiles("templates/result.html")
     if err != nil{
       log.Fatal(err)
     }
     req.ParseForm()
     context := struct{Username string
                     Password string}{req.Form["username"][0],req.Form["password"][0]}
     tmplt.Execute(rw , context)
   }

}

func main() {
	http.HandleFunc("/login", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
