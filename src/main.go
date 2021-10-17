package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", ServeFiles)

	log.Fatal((http.ListenAndServe(":1010", nil)))
}

func ServeFiles(res http.ResponseWriter, req *http.Request) {

	path := req.URL.Path

	if path == "/signup" {
		if req.Method == "GET" {
			http.ServeFile(res, req, "./static/signup.html")
			return
		}
	} else if path == "/signin" {
		if req.Method == "GET" {
			http.ServeFile(res, req, "./static/signin.html")
			return
		}
	} else if path == "/home" {
		if req.Method == "GET" {
			http.ServeFile(res, req, "./static/home.html")
			return
		}
	} else {
		http.ServeFile(res, req, "."+path)
	}
}
