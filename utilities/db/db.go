package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGuser")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDATABASE")
	)
	credStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbname, password, port)
	//fmt.Println(credStr)
	db, err := sql.Open("postgres", credStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}