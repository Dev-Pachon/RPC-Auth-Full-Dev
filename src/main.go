package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/home", homeHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal((http.ListenAndServe(":1010", nil)))
}

func signupHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.ServeFile(res, req, "html/signup.html")
		return
	}
}

func signinHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.ServeFile(res, req, "html/signin.html")
		return
	}
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.ServeFile(res, req, "html/home.html")
		return
	}
}
