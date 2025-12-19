package router

import (
	"Livrable-projet-groupie-tracker/controller"
	"html/template"
	"net/http"
)

func SetupRouter() http.Handler {

	// ----- Gestion du CSS -----
	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	// ----- Template -----
	tmpl := template.Must(template.ParseFiles("templetes/index.html"))

	// ----- Routes -----

	// Page d'accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Recherche
	http.HandleFunc("/search", controller.SearchHandler)
	http.HandleFunc("/dashboard", controller.DashboardHandler)

	// Favoris
	http.HandleFunc("/favorite/add", controller.AddFavoriteHandler)
	http.HandleFunc("/favorite/remove", controller.RemoveFavoriteHandler)

	return http.DefaultServeMux
}
