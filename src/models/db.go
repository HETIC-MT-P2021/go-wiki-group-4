package models

import (
	"database/sql"
	"fmt"

	// Dependancy for mysql
	_ "github.com/go-sql-driver/mysql"
)

var db = sql.DB

// ConnectToDB Set up connection to the mysql db and returns it
func ConnectToDB(host string, dbname string, user string, password string, port string) {
	fmt.Println("Go MySQL ")

	dbParameter := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	// Open up our database connection.
	// set up a database on local machine using phpmyadmin.
	// The database is called gomvc
	tempDB, err := sql.Open("mysql", dbParameter)

	if err != nil {
		fmt.Println("Database connection params error")
		panic(err)
	}

	err = tempDB.Ping()
	if err != nil {
		fmt.Println("Database initialisation error")
		panic(err)
	}

	fmt.Println("Database successfully connected!")

	// defer the close till after the main function has finished
	// executing
	db = tempDB
}
