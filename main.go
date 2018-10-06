package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jamespantalones/jamespants-2018-go/api"
	"github.com/julienschmidt/httprouter"
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

	// port := os.Getenv("PORT")
	// adminRoute := os.Getenv("ADMIN")

	fmt.Println("Starting server...")

	// declare new router
	router := httprouter.New()

	// handle home page
	router.GET("/", api.GetHandler)
	router.GET("/index.html", api.GetHandler)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	// router.POST(adminRoute, api.AdminHandler)

	// httpsRouter := RedirectToHTTPSRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))

}
