package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products []product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func initProducts() {
	bs, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bs, &products); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initProducts()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/products", getProductsHandler)

	// Bind to a port and pass our router in
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
