package main

import (
	"log"
	"net/http"

	"Livrable-projet-groupie-tracker/router"
)

func main() {
	mux := router.New()

	log.Println("Serveur lancé sur http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
