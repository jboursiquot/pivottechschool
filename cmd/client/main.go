package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	res, err := http.Get("http://localhost:8080/products")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var products []product
	err = json.Unmarshal(bs, &products)
	if err != nil {
		log.Fatal(err)
	}

	// for _, p := range products {
	// 	fmt.Printf("%d: %s costs %d\n", p.ID, p.Name, p.Price)
	// }

	spew.Dump(products)
}
