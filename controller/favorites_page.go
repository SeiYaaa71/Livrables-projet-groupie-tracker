package controller

import (
	fonction "Livrable-projet-groupie-tracker/fonctions"
	struct_ "Livrable-projet-groupie-tracker/struct"
	"net/http"
)

func FavoritesPageHandler(w http.ResponseWriter, r *http.Request) {
	// thème
	theme := r.URL.Query().Get("theme")
	themeClass := "theme-classic"
	themeParam := ""
	if theme == "ui" {
		themeClass = "theme-ui"
		themeParam = "theme=ui"
	}

	// favoris -> map
	favIDs := loadFavorites()
	favMap := favoritesToMap(favIDs)
	if favMap == nil {
		favMap = map[int]bool{}
	}

	// récupérer tous les persos en mémoire (via ton cache)
	// IMPORTANT : ApiSearchCharacters("", "", "") doit retourner tous les persos (filtrage vide)
	all := fonction.ApiSearchCharacters("", "", "")

	// filtrer uniquement ceux qui sont en favoris
	results := make([]struct_.CharacterById, 0, len(favIDs))
	for _, c := range all {
		if favMap[c.ID] {
			results = append(results, c)
		}
	}

	data := struct_.FavoritesPageData{
		ThemeClass: themeClass,
		ThemeParam: themeParam,
		Favorites:  favMap,
		Results:    results,
		Total:      len(results),
	}

	renderTemplate(w, "favorites.html", data)
}
