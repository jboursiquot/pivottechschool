package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p *product) validate() error {
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if p.Price == 0 {
		return fmt.Errorf("price is required")
	}
	return nil
}

var products []product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := p.validate(); err != nil {
		log.Printf("product validation error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// set the id to the next available id
	p.ID = len(products) + 1

	// validate the product
	if err := p.validate(); err != nil {
		log.Printf("product validation error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	products = append(products, p)

	w.WriteHeader(http.StatusCreated)
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// convert the id to an int if its present in the request
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("error converting id to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p := lookupProduct(id)
	if p == nil {
		log.Printf("product with id %d not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Printf("error encoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// convert the id to an int if its present in the request
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("error converting id to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p := lookupProduct(id)
	if p == nil {
		log.Printf("product with id %d not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p.ID = id // ensure the id is set to the id in the url

	// validate the product update
	if err := p.validate(); err != nil {
		log.Printf("product validation error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i := range products {
		if id == products[i].ID {
			products[i] = *p
			log.Printf("updated product with id %d", id)
			break
		}
	}
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	// convert the id to an int if its present in the request
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("error converting id to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if p := lookupProduct(id); p == nil {
		log.Printf("product with id %d not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i, prod := range products {
		if prod.ID == products[i].ID {
			products = append(products[:i], products[i+1:]...)
			log.Printf("deleted product with id %d", id)
			break
		}
	}
}

func lookupProduct(id int) *product {
	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func initProducts(filepath string) {
	bs, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bs, &products); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// parse the command line flags for the products filepath
	filePtr := flag.String("file", "products.json", "location of product JSON file")
	flag.Parse()
	initProducts(*filePtr)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/products", createProductHandler).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}", updateProductHandler).Methods(http.MethodPut)
	r.HandleFunc("/products/{id}", deleteProductHandler).Methods(http.MethodDelete)
	r.HandleFunc("/products/{id}", getProductHandler)
	r.HandleFunc("/products", getProductsHandler)

	// Bind to a port and pass our router in
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
