package main

import (
	"fmt"
	"log"
	"net/http"
	"html"	
)

func main(){
	http.Handlefunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.Handlefunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})
	log.Fatal (http.ListenAndServe(":8080", nil))
}