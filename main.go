package main

import (
	"Livrable-projet-groupie-tracker/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	log.Println("Serveur lancé : 🟢 http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
