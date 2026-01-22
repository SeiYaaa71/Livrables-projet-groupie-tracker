package router

import (
	controller "Livrable-projet-groupie-tracker/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Pages
	mux.HandleFunc("/", controller.HomeHandler)
	mux.HandleFunc("/dashboard", controller.DashboardHandler)

	mux.HandleFunc("/search", controller.SearchHandler)
	mux.HandleFunc("/characters", controller.CharactersHandler)
	mux.HandleFunc("/character", controller.CharacterDetailHandler)

	mux.HandleFunc("/favorites", controller.FavoritesPageHandler)
	// Favoris
	mux.HandleFunc("/favorites/add", controller.AddFavoriteHandler)
	mux.HandleFunc("/favorites/remove", controller.RemoveFavoriteHandler)

	// Statics: CSS
	styleServer := http.FileServer(http.Dir("./style"))
	mux.Handle("/style/", http.StripPrefix("/style/", styleServer))

	// Assets (images)
	assetsServer := http.FileServer(http.Dir("./templates/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assetsServer))

	return mux
}
