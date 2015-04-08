package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	host   = os.Getenv("POSTGRESQL_HOST")
	port   = os.Getenv("POSTGRESQL_PORT")
	user   = os.Getenv("POSTGRESQL_USER")
	mainDB = "postgres"
)

func initDB() (*sql.DB, error) {
	connString := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable",
		user, mainDB, host, port)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDatabase(dbname string) (string, *ErrorWithCode) {
	db, err := initDB()
	if err != nil {
		return "", pqError(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
	if err != nil {
		return "", pqError(err)
	}

	_, err = db.Exec(fmt.Sprintf("REVOKE ALL ON DATABASE %s FROM public", dbname))
	if err != nil {
		return "", pqError(err)
	}

	dbString := fmt.Sprintf("http://%s:%s/databases/%s", host, port, dbname)
	return dbString, nil
}