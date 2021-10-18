package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const user string = "root"
const password string = ""
const dbName string = "users"

func ConnectDB() (*sql.DB, error) {

	connectString := fmt.Sprintf("%s:%s@/%s", user, password, dbName)

	db, err := sql.Open("mysql", connectString)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Create(db *sql.DB) error {
	if _, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(20) NOT NULL UNIQUE,
		email VARCHAR(20) NOT NULL UNIQUE,
		password VARCHAR(200) NOT NULL,
		firstname VARCHAR(50) NOT NULL,
		lastname VARCHAR(50) NOT NULL,
		birthdate DATE NOT NULL CHECK (YEAR(birthdate) >= 1820 AND YEAR(birthdate) <= 2016),
		country VARCHAR(50) NOT NULL,
		university VARCHAR(50) NOT NULL
	)
	`); err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func Insert(db *sql.DB, username string, email string, password string, firstname string, lastname string, birthdate string, country string, university string) error {

	hash, errHash := hashPassword(password)

	if errHash != nil {
		return errHash
	}

	DMLSentence := fmt.Sprintf("INSERT INTO users (username, email, password, firstname, lastname, birthdate, country, university) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s')", username, email, hash, firstname, lastname, birthdate, country, university)

	if _, err := db.Exec(DMLSentence); err != nil {
		return err
	}

	return nil
}

func Query(db *sql.DB) (*sql.Rows, error) {

	rows, err := db.Query("SELECT username, email, firstname, lastname, birthdate, country, university FROM users")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func CheckLogin(db *sql.DB, username string, password string) error {
	rows, err := db.Query("SELECT password FROM users WHERE username =?", username)
	if err != nil {
		return err
	}
	var dbPassword string

	rows.Next()

	err = rows.Scan(&dbPassword)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
