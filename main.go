package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// RedirectToHTTPSRouter enforces HTTPS
func RedirectToHTTPSRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		proto := req.Header.Get("x-forwarded-proto")
		if proto == "http" || proto == "HTTP" {
			http.Redirect(res, req, fmt.Sprintf("https://%s%s", req.Host, req.URL), http.StatusPermanentRedirect)
			return
		}

		next.ServeHTTP(res, req)

	})
}

// Main
func main() {

	port := os.Getenv("PORT")
	adminRoute := os.Getenv("ADMIN")

	fmt.Println("Starting server...")

	// declare new router
	r := mux.NewRouter()

	// handle home page
	r.HandleFunc("/", GetHandler)
	r.HandleFunc("/index.html", GetHandler)
	r.HandleFunc(adminRoute, AdminHandler).Methods("POST")

	// handle static files
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))

	// match all routes starting with /static/
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	httpsRouter := RedirectToHTTPSRouter(r)

	log.Fatal(http.ListenAndServe(":"+port, httpsRouter))

}
