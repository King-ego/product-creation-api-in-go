package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres_go"
	port     = 5432
	user     = "ted"
	password = "1234"
	dbname   = "pd_users"
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
