package controller

import (
	"Livrable-projet-groupie-tracker/fonctions"
	"encoding/json"
	"html/template"
	struct_ "Livrable-projet-groupie-tracker/struct"
	"net/http"
	"os"
	"strconv"
)

var favoritesFile = "favorites.json"

//fonction pour charger les pages html
func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	tmpl.Execute(w, data)
}

// Charger favoris depuis JSON
func loadFavorites() []int {
	data, err := os.ReadFile(favoritesFile)
	if err != nil {
		return []int{}
	}

	var fav []int
	json.Unmarshal(data, &fav)
	return fav
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	theme := r.URL.Query().Get("theme")
	
	// Utilisation de struct_.SearchPageData
	data := struct_.SearchPageData{
		ThemeClass: "",
		ThemeParam: "",
	}

	if theme == "ui" {
		data.ThemeClass = "ui-theme"
		data.ThemeParam = "?theme=ui"
	}

	tmpl, err := template.ParseFiles("templetes/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Sauvegarder favoris
func saveFavorites(fav []int) {
	data, _ := json.Marshal(fav)
	os.WriteFile(favoritesFile, data, 0644)
}

// Page d’accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templetes/index.html")
	tmpl.Execute(w, nil)
	fonction.ApiGet("characters", []string{})
	data := fonction.Data
	renderTemplate(w, "index.html", data)
}

func FilterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templetes/search.html")
	tmpl.Execute(w, nil)
	var detail string = "characters/1"
	fonction.ApiGet(detail, []string{})
	data := " "
	renderTemplate(w, "search.html", data)
}

// Recherche + filtres + pagination
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}

	results := []struct_.Characters{}


	// Pagination par 10
	start := (page - 1) * 10
	end := start + 10

	if start > len(results) {
		start = len(results)
	}
	if end > len(results) {
		end = len(results)
	}

	paged := results[start:end]

	response, _ := json.Marshal(paged)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Ajouter un favori
func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	fav := loadFavorites()
	fav = append(fav, id)
	saveFavorites(fav)

	w.Write([]byte("Ajouté aux favoris"))
}

// Supprimer un favori
func RemoveFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	fav := loadFavorites()
	newFav := []int{}
	for _, v := range fav {
		if v != id {
			newFav = append(newFav, v)
		}
	}

	saveFavorites(newFav)
	w.Write([]byte("Supprimé des favoris"))
}
