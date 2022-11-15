package main

import (
	"database/sql"
	"log"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "zip.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, code, city, state FROM places")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type place struct {
		id    int
		code  string
		city  string
		state string
	}

	var places []place

	for rows.Next() {
		var p place
		err = rows.Scan(&p.id, &p.code, &p.city, &p.state)
		if err != nil {
			log.Fatal(err)
		}
		places = append(places, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(places)

}
