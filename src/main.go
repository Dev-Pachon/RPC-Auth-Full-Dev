package main

import (
	"RPC-AUTH-FULL-DEV/src/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
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
	Username   string
	Email      string
	Password   string
	Firstname  string
	Lastname   string
	Birthdate  string
	Country    string
	University string
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
				responseData.Result = "Nothing is ok"
				responseData.Content = "Error parsing data: " + err.Error()
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			db, err := database.ConnectDB()

			if err != nil {
				fmt.Println("Cannot connect to database")
				var responseData VerifyUserOutput
				responseData.Result = "Nothing is ok"
				responseData.Content = "Cannot connect to database"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			defer db.Close()

			err = database.Insert(db, data.Username, data.Email, data.Password, data.Firstname, data.Lastname, data.Birthdate, data.Country, data.University)

			if err != nil {
				fmt.Println("Cannot insert the element to the database")
				fmt.Println(err)
				var responseData VerifyUserOutput
				responseData.Result = "Nothing is ok"
				responseData.Content = "The username or email is already in use"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			var responseData VerifyUserOutput
			responseData.Result = "ok"
			res.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(res).Encode(responseData)

			sendEmail(data.Email)

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
				responseData.Result = "Nothing is ok"
				responseData.Content = "Error parsing data: " + err.Error()
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			db, err := database.ConnectDB()

			if err != nil {
				fmt.Println("Cannot connect to database")
				var responseData VerifyUserOutput
				responseData.Result = "Nothing is ok"
				responseData.Content = "Cannot connect to database"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			defer db.Close()

			token, err := database.CheckLogin(db, data.Username, data.Password)

			if err != nil {
				fmt.Println("The username or the password is incorrect")
				fmt.Println(err)
				var responseData VerifyUserOutput
				responseData.Result = "Nothing is ok"
				responseData.Content = "The username or the password is incorrect"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			var responseData VerifyUserOutput
			responseData.Result = "ok"
			responseData.Content = data.Username
			responseData.SessionLogin = token
			res.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(res).Encode(responseData)
			return
		}
	} else if strings.HasPrefix(path, "/home") {

		if req.Method == "GET" {
			pathWithoutPrefix := strings.TrimPrefix(path, "/home-")

			pathArr := strings.Split(pathWithoutPrefix, "-")

			if len(pathArr) != 2 {
				return
			}

			db, err := database.ConnectDB()

			if err != nil {
				fmt.Println("Cannot connect to database")
				var responseData VerifyUserOutput
				responseData.Result = "Nothing is ok"
				responseData.Content = "Cannot connect to database"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			err = database.CheckToken(db, pathArr[0], pathArr[1])

			if err != nil {
				fmt.Println(err)
				return
			}

			http.ServeFile(res, req, "./static/home.html")
			return
		}

		if req.Method == "SELECT" {

			db, err := database.ConnectDB()

			if err != nil {
				fmt.Println("Cannot connect to database")
				var responseData VerifyUserOutput
				responseData.Result = "Nothing is ok"
				responseData.Content = "Cannot connect to database"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			rows, err := database.Query(db)

			if err != nil {
				fmt.Println("Could not retrieve users from the database")
				fmt.Println(err)
				var responseData VerifyUserOutput
				responseData.Result = "nok"
				responseData.Content = "Could not retrieve users from the database"
				res.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(res).Encode(responseData)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(res).Encode(rows)
			return
		}

	} else {
		http.ServeFile(res, req, "."+path)
	}
}

func sendEmail(email string) {
	// Sender data.
	from := "t3stem4ilauth@gmail.com"
	password := "ytrewq654321"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(`
	Dear user, your account was created correctly.

	Visit our web site: https://redprogramacioncompetitiva.com
	`)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
