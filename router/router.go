package router

import (
	controller "Livrable-projet-groupie-tracker/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.SearchHandler)

	return mux
}
