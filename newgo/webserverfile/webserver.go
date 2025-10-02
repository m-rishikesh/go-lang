package webserverfile

import (
	"net/http"
)

func Webserver() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the homepage by go"))
}
