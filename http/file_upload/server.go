package main

import (
  "os"
  "io"
  "fmt"
  "log"
  "net/http"
  "html/template"
)

func file_handler(rw http.ResponseWriter , req *http.Request){
  if req.Method == "GET"{
    tmplt , err := template.ParseFiles("index.html")
    if err != nil{
      log.Fatal(err)
    }
    tmplt.Execute(rw , nil)
  }else{
    req.ParseMultipartForm(32 << 20)
    file , handler , err := req.FormFile("myfile")
    if err != nil{
      log.Fatal(err)
    }
    defer file.Close()
    f , err := os.OpenFile("./uploaded/" + handler.Filename , os.O_WRONLY|os.O_CREATE , 0666)
    if err != nil{
      log.Fatal(err)
    }
    defer f.Close()
    io.Copy(f , file)
    fmt.Fprintf(rw , "File Uploaded")
  }
}

func main(){
  http.HandleFunc("/upload" , file_handler)
  log.Println("Server starting... ... ...")
  err := http.ListenAndServe(":8080" , nil)
  if err != nil{
    log.Fatal(err)
  }
}
