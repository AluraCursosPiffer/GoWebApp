package main

import (
	"database/sql"
	"net/http"

	"github.com/AluraCursosPiffer/GoWebApp/database"
	"github.com/AluraCursosPiffer/GoWebApp/routes"
	_ "github.com/lib/pq"
)

func main() {
	var Db *sql.DB = database.DatabaseConection()

	routes.Routes()
	http.ListenAndServe(":8000", nil)

	Db.Close()
}
