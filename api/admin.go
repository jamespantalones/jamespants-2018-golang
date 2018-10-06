// Admin Handler

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Row is the final row type
type Row struct {
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	URL         string  `json:"url"`
	Year        float64 `json:"year"`
	Type        string  `json:"type"`
}

// Final is returned data
type Final struct {
	Updated int   `json:"updated"`
	Rows    []Row `json:"rows"`
}

// Request is incoming
type Request struct {
	Body struct {
		Updated int
		Rows    [][]interface{}
	}
}

// AdminHandler deals with incoming signal from Google Apps Scrip
func AdminHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// read in body as slice of byte
	body, errr := ioutil.ReadAll(r.Body)

	// handle error
	if errr != nil {
		log.Printf("Error reading body: %v", errr)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	var req Request

	err := json.Unmarshal(body, &req)

	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		http.Error(w, "Can't unmarshal JSON", http.StatusBadRequest)
	}

	final := Final{
		Updated: req.Body.Updated,
	}

	// make keys
	var keys []string

	for index, row := range req.Body.Rows {
		// make a new row
		var r Row

		fmt.Println(row, keys)

		for i, col := range row {
			if index == 0 {
				keys = append(keys, col.(string))
			} else {
				switch keys[i] {
				case "Title":
					r.Title = col.(string)
				case "Description":
					r.Description = col.(string)
				case "URL":
					r.URL = col.(string)
				case "Year":
					r.Year = col.(float64)

				case "Type":
					r.Type = col.(string)
				default:
					fmt.Println("Missing")
				}
			}
		}

		if index != 0 {
			final.Rows = append(final.Rows, r)
		}
	}

	// response
	out, err := json.Marshal(final)

	if err != nil {
		panic(err)
	}

	// write file
	werr := ioutil.WriteFile("./static/data.json", out, 0644)

	if werr != nil {
		fmt.Println(werr)
	}

	fmt.Fprintf(w, "Ok")
}
