package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{{
		Name:        "",
		Description: "",
		Price:       0,
		Quantity:    0,
	}}

	temp := template.Must(template.ParseGlob("templates/*.html"))
	temp.ExecuteTemplate(w, "Index", products)
}
