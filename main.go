package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Main
func main() {

	port := os.Getenv("PORT")
	adminRoute := os.Getenv("ADMIN")

	fmt.Println("Starting server...")

	// declare new router
	r := mux.NewRouter()

	// handle home page
	r.HandleFunc("/", GetHandler)
	r.HandleFunc(adminRoute, AdminHandler).Methods("POST")

	// handle static files
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))

	// match all routes starting with /static/
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	http.ListenAndServe(":"+port, r)

}
