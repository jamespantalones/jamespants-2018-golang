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

type MusicItem struct {
	Artist string
	Album  string
	Year   int
}

type MusicItems struct {
	Items []*MusicItem
}

// TechItem is individual technology items
type TechItem struct {
	Name string
}

// TechItems is Colleciton of tech items
type TechItems struct {
	Items []*TechItem
}

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
	PageTitle  string
	Items      []*Item
	TechItems  *TechItems
	MusicItems *MusicItems
}

// GetHandler displays all info
func GetHandler(w http.ResponseWriter, r *http.Request) {

	//var page PageData
	var items []*Item
	var techItems *TechItems
	var musicItems *MusicItems

	// Get all items in data foler
	matches, err := filepath.Glob("./data/portfolio/*.toml")

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

	// read tech
	te, err := ioutil.ReadFile("./data/tech/tech.toml")

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Convert tech to string
	techStr := string(te)

	// Decode tech as toml
	if _, err := toml.Decode(techStr, &techItems); err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	// read music
	me, err := ioutil.ReadFile("./data/music/music.toml")

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	musicStr := string(me)

	if _, err := toml.Decode(musicStr, &musicItems); err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Parse template
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/style.html"))

	// sort slice
	sort.Slice(items, func(i, j int) bool {
		return items[i].Year > items[j].Year
	})

	// create page data for template
	data := PageData{
		PageTitle:  "James Pants",
		Items:      items,
		TechItems:  techItems,
		MusicItems: musicItems,
	}

	// run template
	t.ExecuteTemplate(w, "layout", data)

}
