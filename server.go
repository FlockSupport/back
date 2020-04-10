package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)

var router *chi.Mux
var db *sql.DB

func routers() *chi.Mux {
	router.Get("/posts", AllPosts)
	return router
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "graphql_demo"
)

func init() {
	var err error
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query(`SELECT * from graphql`)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var quantity int
		var id int
		var name string
		err = result.Scan(&id, &name, &quantity)
		fmt.Printf("%v%s", quantity, name)
		panic(err)

	}
	err = result.Err()
}

func main() {
	defer db.Close()
	routers()
	http.ListenAndServe(":8005", router)
}
