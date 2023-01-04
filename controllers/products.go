package controllers

import (
	"database/sql"
	"net/http"
	"text/template"

	"github.com/AluraCursosPiffer/GoWebApp/structs"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var Db *sql.DB

	products, err := Db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := structs.Product{}
	productsList := []structs.Product{}

	for products.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = products.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		productsList = append(productsList, p)
	}

	temp := template.Must(template.ParseGlob("templates/*.html"))
	temp.ExecuteTemplate(w, "Index", products)
}

