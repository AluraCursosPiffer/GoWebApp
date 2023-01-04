package routes

import (
	"net/http"

	"github.com/AluraCursosPiffer/GoWebApp/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
}
