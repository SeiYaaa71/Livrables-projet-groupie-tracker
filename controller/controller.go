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

//
// ========================
// UTILS
// ========================
//

// Render template centralisé
func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Charger favoris
func loadFavorites() []int {
	data, err := os.ReadFile(favoritesFile)
	if err != nil {
		return []int{}
	}

	var fav []int
	json.Unmarshal(data, &fav)
	return fav
}

// Sauvegarder favoris
func saveFavorites(fav []int) {
	data, _ := json.Marshal(fav)
	os.WriteFile(favoritesFile, data, 0644)
}

//
// ========================
// HANDLERS PAGES
// ========================
//

// Page d’accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fonction.ApiGet("characters", []string{})

	renderTemplate(w, "index.html", fonction.Data)
}

// Page dashboard
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	theme := r.URL.Query().Get("theme")

	data := struct_.SearchPageData{
		ThemeClass: "",
		ThemeParam: "",
	}

	if theme == "ui" {
		data.ThemeClass = "ui-theme"
		data.ThemeParam = "?theme=ui"
	}

	renderTemplate(w, "dashboard.html", data)
}

// Page recherche simple
func FilterPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "search.html", nil)
}

// Page recherche avancée
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	race := r.URL.Query().Get("race")
	affiliation := r.URL.Query().Get("affiliation")

	results := fonction.ApiSearchCharacters(name, race, affiliation)

	data := struct_.SearchResultsData{
		Query:       name,
		Race:        race,
		Affiliation: affiliation,
		Results:     results,
	}

	renderTemplate(w, "RealSearch.html", data)
}


//
// ========================
// FAVORIS
// ========================
//

// Ajouter un favori
func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	fav := loadFavorites()
	fav = append(fav, id)
	saveFavorites(fav)

	w.Write([]byte("Ajouté aux favoris"))
}

// Supprimer un favori
func RemoveFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

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

func CharactersHandler(w http.ResponseWriter, r *http.Request) {
	fonction.ApiGet("characters", []string{})

	renderTemplate(w, "characters.html", fonction.Data)
}
