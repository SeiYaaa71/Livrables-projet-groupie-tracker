package router

import (
	"Livrable-projet-groupie-tracker/controller"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.HomeHandler).Methods("GET")
	r.HandleFunc("/search", controller.SearchHandler).Methods("GET")
	r.HandleFunc("/favorite/add", controller.AddFavoriteHandler).Methods("POST")
	r.HandleFunc("/favorite/remove", controller.RemoveFavoriteHandler).Methods("POST")

	return r
}
