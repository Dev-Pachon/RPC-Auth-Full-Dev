package main

import (
	"RPC-AUTH-FULL-DEV/src/database"
	"log"
	"net/http"
)

func main() {
	db, err := database.ConnectDB()

	if err != nil {
		panic(err)
	}

	err = database.Create(db)

	if err != nil {
		panic(err)
	}

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
		if req.Method == "POST" {
			//TODO
			return
		}
	} else if path == "/signin" {
		if req.Method == "GET" {
			http.ServeFile(res, req, "./static/signin.html")
			return
		}
		if req.Method == "POST" {
			//TODO
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
