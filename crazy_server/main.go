package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items []Item

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	items = append(items, item)

	fmt.Fprintf(w, "Id: %v Created Item: %s with Price: %.2f", item.Id, item.Name, item.Price)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	for index, item := range items {
		if item.Id == params["id"] {
			items = append(items[:index], items[index+1:]...)
			json.NewEncoder(w).Encode(items)

			return
		}
	}
}

func patchItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	for index, item := range items {
		if item.Id == params["id"] {
			items = append(items[:index], items[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&newItem)

			newItem.Id = params["id"]
			newItem.Price = 1000
			items = append(items, newItem)
			json.NewEncoder(w).Encode(newItem)
		}
	}
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create-item-using-delete", createItemHandler).Methods("DELETE")
	r.HandleFunc("/get-all", getAllHandler).Methods("GET")
	r.HandleFunc("/delete-item-using-post/{id}", deleteItemHandler).Methods("POST")
	r.HandleFunc("/patch-item-using-get/{id}", patchItemHandler).Methods("GET")
	fmt.Println(`
 _____     _    _                        _                 _                  _      _   ___ ___ ___ ___
|  _  |___|_|  |_|___    ___ _ _ ___ ___|_|___ ___    ____| |_    ___ ___ ___| |_   |_| | . |   | . |   |
|     | . | |  | |_ -|  |  _| | |   |   | |   | . |  | .' |  _|  | . | . |  _|  _|   _  | . | | | . | | |
|__|__|  _|_|  |_|___|  |_| |___|_|_|_|_|_|_|_|_  |  |__,_|_|    |  _|___|_| |_|    |_| |___|___|___|___|
      |_|                                     |___|              |_|
	`)
	log.Fatal(http.ListenAndServe(":8080", r))
}
