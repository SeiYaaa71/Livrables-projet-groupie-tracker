package controller

import (
	fonction "Livrable-projet-groupie-tracker/fonctions"
	struct_ "Livrable-projet-groupie-tracker/struct"
	"bytes"
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

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buf.Bytes())
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
	themeClass, themeParam := ResolveTheme(r)

	data := struct_.HomePageData{
		ThemeClass: themeClass,
		ThemeParam: themeParam,
	}

	renderTemplate(w, "index.html", data)
}

func CharactersHandler(w http.ResponseWriter, r *http.Request) {
	themeClass, themeParam := ResolveTheme(r)

	items := fonction.GetCharactersCached() // []CharacterById

	data := struct {
		ThemeClass string
		ThemeParam string
		Items      []struct_.CharacterById
	}{
		ThemeClass: themeClass,
		ThemeParam: themeParam,
		Items:      items,
	}

	renderTemplate(w, "characters.html", data)
}

func CharacterDetailHandler(w http.ResponseWriter, r *http.Request) {
	themeClass, themeParam := ResolveTheme(r)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	c, ok := fonction.GetCharacterByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	favIDs := loadFavorites()
	data := map[string]interface{}{
		"ThemeClass": themeClass,
		"ThemeParam": themeParam,
		"Character":  c,
		"IsFav":      favoritesToMap(favIDs)[id],
	}

	renderTemplate(w, "character_detail.html", data)
}

// Page dashboard
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	PersistThemeCookie(w, r)
	themeClass, themeParam := ResolveTheme(r)

	data := struct_.SearchPageData{
		ThemeClass: themeClass,
		ThemeParam: themeParam,
	}

	renderTemplate(w, "dashboard.html", data)
}

// Page recherche simple
func FilterPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "search.html", nil)
}

//
// ========================
// FAVORIS
// ========================
//

// Ajouter un favori
func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	fav := loadFavorites()
	for _, v := range fav {
		if v == id {
			http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
			return
		}
	}

	fav = append(fav, id)
	saveFavorites(fav)

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func RemoveFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
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
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}
