package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Starting server...")

	// connect to db

	connString := "dbname=jamespants sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	fmt.Println("Connected to DB")

	// declare new router
	r := mux.NewRouter()

	r.HandleFunc("/", GetHandler)

	// handle static files
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))

	// match all routes starting with /static/
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	http.ListenAndServe(":8000", r)

}
