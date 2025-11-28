package main

import (
	"log"
	"net/http"
	"dbzapp/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Serveur en cours sur : http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
