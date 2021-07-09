package main

import (
  "net/http"
  "fmt"
)

func main() {    
     fmt.Println("Server running in http://localhost:8000")
     http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
     http.ListenAndServe(":8000", nil)
}

