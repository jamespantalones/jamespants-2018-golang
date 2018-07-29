package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

// JSONData is main JSON structure
type JSONData struct {
	Updated int
	Rows    []*Item
}

// GetHandler displays all info
func GetHandler(w http.ResponseWriter, r *http.Request) {

	var d JSONData

	// get json file
	dat, err := ioutil.ReadFile("./static/data.json")

	if err != nil {
		fmt.Println(fmt.Errorf("Error fetching data file: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	err2 := json.Unmarshal(dat, &d)

	if err2 != nil {
		fmt.Println(fmt.Errorf("Error decoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Parse template
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/style.html"))

	// sort slice
	// sort.Slice(items, func(i, j int) bool {
	// 	return items[i].Year > items[j].Year
	// })

	// create page data for template
	data := PageData{
		PageTitle: "James Pants",
		Items:     d.Rows,
	}

	// run template
	t.ExecuteTemplate(w, "layout", data)

}
