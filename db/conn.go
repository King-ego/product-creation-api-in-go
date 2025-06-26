package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, error := sql.Open("postgres", psqlInfo)
	if error != nil {
		panic(error)
	}

	error = db.Ping()

	if error != nil {
		panic(error)

	}

	fmt.Println("connected to " + dbname)

	return db, nil
}
