package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Name        string `json:"item"`
	Description string `json:"description"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	items, err := store.GetItems()

	fmt.Println(items)

	itemsListBytes, err := json.Marshal(items)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(itemsListBytes)

	//fmt.Fprintf(w, "Hello World!")

}
