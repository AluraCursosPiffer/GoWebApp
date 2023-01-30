package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DatabaseConection() *sql.DB {
	conectionString := "user=postgres dbname=aluna_loja password=Gui090900! sslmode=disable"
	Db, err := sql.Open("postgres", conectionString)
	if err != nil {
		log.Fatal(err)
	}

	return Db
}
