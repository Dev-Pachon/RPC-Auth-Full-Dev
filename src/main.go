package main

import (
	"RPC-AUTH-FULL-DEV/src/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := database.ConnectDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = database.Create(db)

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", ServeFiles)
	fmt.Println("Serving at: localhost:1010/")
	log.Fatal((http.ListenAndServe(":1010", nil)))
}

type user struct {
	Username  string
	Password  string
	Firstname string
	Lastname  string
	Birthdate string
}

type VerifyUserInput struct {
	Username string
	Password string
}

type VerifyUserOutput struct {
	Result       string
	Content      string
	SessionLogin string
}

func ServeFiles(res http.ResponseWriter, req *http.Request) {

	path := req.URL.Path

	if path == "/signup" {
		if req.Method == "GET" {
			http.ServeFile(res, req, "./static/signup.html")
			return
		}
		if req.Method == "POST" {
			var data user
			err := json.NewDecoder(req.Body).Decode(&data)

			if err != nil {
				fmt.Println("Error parsing data: " + err.Error())
				var responseData VerifyUserOutput
				responseData.Result = "nok"
				responseData.Content = "Error parsing data: " + err.Error()
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			db, err := database.ConnectDB()

			if err != nil {
				fmt.Println("Cannot connect to database")
				var responseData VerifyUserOutput
				responseData.Result = "nok"
				responseData.Content = "Cannot connect to database"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			defer db.Close()

			err = database.Insert(db, data.Username, data.Password, data.Firstname, data.Lastname, data.Birthdate)

			if err != nil {
				fmt.Println("Cannot insert the element to the database")
				fmt.Println(err)
				var responseData VerifyUserOutput
				responseData.Result = "nok"
				responseData.Content = "Cannot insert the element to the database"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			var responseData VerifyUserOutput
			responseData.Result = "ok"
			res.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(res).Encode(responseData)
			return
		}
	} else if path == "/signin" {
		if req.Method == "GET" {
			http.ServeFile(res, req, "./static/signin.html")
			return
		}
		if req.Method == "POST" {
			var data VerifyUserInput
			err := json.NewDecoder(req.Body).Decode(&data)

			if err != nil {
				fmt.Println("Error parsing data: " + err.Error())
				var responseData VerifyUserOutput
				responseData.Result = "nok"
				responseData.Content = "Error parsing data: " + err.Error()
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			db, err := database.ConnectDB()

			if err != nil {
				fmt.Println("Cannot connect to database")
				var responseData VerifyUserOutput
				responseData.Result = "nok"
				responseData.Content = "Cannot connect to database"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			defer db.Close()

			err = database.CheckLogin(db, data.Username, data.Password)

			if err != nil {
				fmt.Println("The username or the password is incorrect")
				fmt.Println(err)
				var responseData VerifyUserOutput
				responseData.Result = "nok"
				responseData.Content = "The username or the password is incorrect"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			var responseData VerifyUserOutput
			responseData.Result = "ok"
			res.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(res).Encode(responseData)
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
