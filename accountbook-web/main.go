package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./accountbook.db")

	if err != nil {
		log.Fatal(err)
	}

	ab := NewAccountBook(db)

	if err := ab.CreateTable(); err != nil {
		log.Fatal(err)
	}

	hs := NewHandlers(ab)
	http.HandleFunc("/", hs.ListHandler)
	http.HandleFunc("/save", hs.SaveHandler)
	http.HandleFunc("/summary", hs.SummaryHandler)

	println("http://localhost:8080 で起動中...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}