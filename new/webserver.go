package main

import (
	"net/http"
)

func webserver() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the homepage by go"))
}
