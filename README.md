# RPC-Auth-Full-Dev
Authentication system for RPC web page using javascript and go.

## How to use

You must have installed MySQL and Go, then inside of the file database.go in database package replace in the lines 14, 17 and 20 with your own credentials.

```
//Username in your mysql database
const user string = "user"

//Password in your mysql database
const password string = "password"

//Name of the database
const dbName string = "dbName"
```


## How to run

To run the program, open your powershell or cmd in the folder src of the project, and type this command:

```
go run .
```

Finally, open your browser and go to http://localhost:1010/signin
