package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sort"

	"github.com/BurntSushi/toml"
)

// Item is Individual Item, whether web or music project
type Item struct {
	Description string
	Title       string
	URL         template.URL
	Year        int
	Type        string
}

// PageData is each HTML template struct
type PageData struct {
	PageTitle string
	Items     []*Item
}

// GetHandler displays all info
func GetHandler(w http.ResponseWriter, r *http.Request) {

	//var page PageData
	var items []*Item

	// Get all items in data foler
	matches, err := filepath.Glob("./data/*.toml")

	// Handle error
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	// For each toml file
	for _, file := range matches {
		// read the file as bytes
		b, err := ioutil.ReadFile(file)

		// Handle error
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
		}

		// Create an Item
		var itemX Item

		// Coerce to string
		str := string(b)

		// Decode TOML
		if _, err := toml.Decode(str, &itemX); err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
		}

		// append item
		items = append(items, &itemX)

	}

	// Parse template
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/style.html"))

	// sort slice
	sort.Slice(items, func(i, j int) bool {
		return items[i].Year > items[j].Year
	})
	// create page data for template
	data := PageData{
		PageTitle: "James Pants",
		Items:     items,
	}

	// run template
	t.ExecuteTemplate(w, "layout", data)

}
