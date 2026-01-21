package router

import (
	controller "Livrable-projet-groupie-tracker/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Pages
	mux.HandleFunc("/", controller.HomeHandler)
	mux.HandleFunc("/search", controller.SearchHandler)
	mux.HandleFunc("/dashboard", controller.DashboardHandler)
	mux.HandleFunc("/characters", controller.CharactersHandler)

	// Fichiers statiques (CSS)
	fileServer := http.FileServer(http.Dir("./style"))
	mux.Handle("/style/", http.StripPrefix("/style/", fileServer))

	return mux
}
