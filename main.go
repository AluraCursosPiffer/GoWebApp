package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

var Db *sql.DB

func databaseConection() *sql.DB {
	conectionString := "user=postgres dbname=aluna_loja password=Gui090900! sslmode=disable"
	db, err := sql.Open("postgres", conectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	Db := databaseConection()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

	Db.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	products, err := Db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	productsList := []Product{}

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
