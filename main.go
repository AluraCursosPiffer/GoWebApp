package main

import (
	"net/http"

	"github.com/AluraCursosPiffer/GoWebApp.git/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
