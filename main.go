package main

import (
	"log"
	"net/http"

	fonction "Livrable-projet-groupie-tracker/fonctions"
	"Livrable-projet-groupie-tracker/router"
)

func main() {
	// Charger le cache une fois
	if err := fonction.LoadCharactersCache(); err != nil {
		log.Fatal("Erreur chargement cache personnages :", err)
	}

	mux := router.New()

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
