package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/julienschmidt/httprouter"
)

// Item is Individual Item, whether web or music project
type Item struct {
	Description string
	Title       string
	URL         template.URL
	Year        float64
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
func GetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var d JSONData

	// get json file
	dat, err := ioutil.ReadFile("./static/data.json")

	if err != nil {
		fmt.Println(fmt.Errorf("Error fetching data file: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	err2 := json.Unmarshal(dat, &d)

	if err2 != nil {
		fmt.Println(err2)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Parse template
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/style.html"))

	items := d.Rows

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
