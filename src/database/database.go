package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const User string = "user"
const Password string = "password"
const DBName string = "dbname"

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", `${User}:${Password}@/${DBName}`)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func create(db *sql.DB) error {
	if _, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(20) NOT NULL UNIQUE,
		password VARCHAR(20) NOT NULL,
		firstname VARCHAR(50) NOT NULL,
		lastname VARCHAR(50) NOT NULL,
		birthdate DATE NOT NULL
	)
	`); err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func insert(db *sql.DB, username string, password string, firstname string, lastname string, birthdate string) error {

	hash, errHash := HashPassword(password)

	if errHash != nil {
		return errHash
	}

	DMLSentence := fmt.Sprintf("INSERT INTO users (username, password, firstname, lastname, birthdate) VALUES (%s,%s,%s,%s,%s)", username, hash, firstname, lastname, birthdate)

	if _, err := db.Exec(DMLSentence); err != nil {
		return err
	}

	return nil
}

func query(db *sql.DB) (*sql.Rows, error) {

	rows, err := db.Query("SELECT username, firstname, lastname, birthdate FROM users")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func checkLogin(db *sql.DB, username string, password string) error {
	row, err := db.Query("SELECT password FROM users WHERE username =?", username)
	if err != nil {
		return err
	}
	var dbPassword string

	err = row.Scan(&dbPassword)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
