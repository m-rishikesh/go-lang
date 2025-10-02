package main

import (
	"fmt"
	"net/http"
)

func Webserver() {
	http.HandleFunc("/", handler1)
	fmt.Println("running server at 8080....")
	http.ListenAndServe(":8080", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the homepage by go"))
}
