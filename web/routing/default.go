package main

import (
	"fmt"
	"net/http"
)

func main() {
	  http.HandleFunc("/dog",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"Hello dog")
	  })
       http.HandleFunc("/cat",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"Hello cat")
	  })
	  http.ListenAndServe(":8080",nil)

}