package main

import (
	router "Livrable-projet-groupie-tracker/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.New()

	fileServer := http.FileServer(http.Dir("./template"))

	r.Handle("/template/", http.StripPrefix("/template/", fileServer))

	fmt.Println("serveur démare sur http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
