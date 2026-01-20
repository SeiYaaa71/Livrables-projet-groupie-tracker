package controller

import (
	fonction "Livrable-projet-groupie-tracker/fonctions"
	struct_ "Livrable-projet-groupie-tracker/struct"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

var favoritesFile = "favorites.json"

// fonction pour charger les pages html
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

	tmpl, err := template.ParseFiles("templates/dashboard.html")
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
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
	fonction.ApiGet("characters", []string{})
	data := fonction.Data
	renderTemplate(w, "index.html", data)
}

func FilterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/search.html")
	tmpl.Execute(w, nil)
	var detail string = "characters/1"
	fonction.ApiGet(detail, []string{})
	data := " "
	renderTemplate(w, "search.html", data)
}

// Recherche + filtres + pagination
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if r.Method == http.MethodGet {
		Id := r.FormValue("id")

		if Id != " " {
			fonction.ApiGet("characters/"+Id, []string{})
			data["character"] = fonction.Data
		}
	}
	data["Characters"] = nil
	renderTemplate(w, "RealSearch.html", data)

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
