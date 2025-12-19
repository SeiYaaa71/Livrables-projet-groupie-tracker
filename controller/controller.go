package controller

import (
	"Livrable-projet-groupie-tracker/fonctions"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Character struct {
	ID         int
	Name       string
	Race       string
	Affiliation string
	FirstYear  int
}

var characterDB = []Character{
	{1, "Goku", "Saiyan", "Z-Fighters", 1984},
	{2, "Vegeta", "Saiyan", "Z-Fighters", 1988},
	{3, "Gohan", "Saiyan/Humain", "Z-Fighters", 1989},
	{4, "Piccolo", "Namek", "Z-Fighters", 1984},
	{5, "Freezer", "Alien", "Armée de Freezer", 1989},
	{6, "Cell", "Bio-Android", "Armée du Ruban Rouge (RR)", 1992},
	{7, "Majin Boo", "Démon", "Aucun", 1994},
	{8, "Trunks", "Saiyan/Humain", "Z-Fighters", 1991},
	{9, "Bulma", "Humain", "Z-Fighters", 1984},
	{10, "Krillin", "Humain", "Z-Fighters", 1984},
	// Tu peux en ajouter autant que tu veux
}

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
	name := r.URL.Query().Get("name")
	race := r.URL.Query().Get("race")
	affiliation := r.URL.Query().Get("affiliation")
	year := r.URL.Query().Get("year")

	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}

	results := []Character{}

	for _, c := range characterDB {

		if name != "" && !strings.Contains(strings.ToLower(c.Name), strings.ToLower(name)) {
			continue
		}

		if race != "" && c.Race != race {
			continue
		}

		if affiliation != "" && c.Affiliation != affiliation {
			continue
		}

		if year != "" {
			y, _ := strconv.Atoi(year)
			if c.FirstYear != y {
				continue
			}
		}

		results = append(results, c)
	}

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

// DashboardHandler affiche la page avec les petits boutons et la barre de recherche
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
    theme := r.URL.Query().Get("theme")
    data := SearchHandler{ThemeClass: ""}
    
    if theme == "ui" {
        data.ThemeClass = "ui-theme"
    }

    tmpl, err := template.ParseFiles("templetes/dashboard.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, data)
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
