package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Item is Individual Item, whether web or music project
type Item struct {
	Name        string `json:"item"`
	Description string `json:"description"`
	URL         string `json:"link"`
}

// PageData is each HTML template struct
type PageData struct {
	PageTitle string
	Items     []*Item
}

// GetHandler displays all info
func GetHandler(w http.ResponseWriter, r *http.Request) {

	// get all items in DB
	items, err := store.GetItems()

	// If we want to display as JSON
	//itemsListBytes, err := json.Marshal(items)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Parse template
	t := template.Must(template.ParseFiles("templates/layout.html"))

	// create page data for template
	data := PageData{
		PageTitle: "James Pants",
		Items:     items,
	}

	// run template
	t.ExecuteTemplate(w, "layout", data)

}
