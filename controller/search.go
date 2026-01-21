package controller

import (
	fonction "Livrable-projet-groupie-tracker/fonctions"
	struct_ "Livrable-projet-groupie-tracker/struct"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Convertit la liste des favoris (slice) en map pour le template
func favoritesToMap(ids []int) map[int]bool {
	m := make(map[int]bool, len(ids))
	for _, id := range ids {
		m[id] = true
	}
	return m
}

// Convertit "60.000.000" / "unknown" en int pour trier
func kiToInt(s string) int {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	if b.Len() == 0 {
		return 0
	}
	n, err := strconv.Atoi(b.String())
	if err != nil {
		return 0
	}
	return n
}

// Tri des personnages selon sortBy + order
// sortBy: "name" | "ki" | "maxKi"
// order:  "asc" | "desc"
func sortCharacters(chars []struct_.CharacterById, sortBy, order string) {
	sortBy = strings.TrimSpace(strings.ToLower(sortBy))
	order = strings.TrimSpace(strings.ToLower(order))
	desc := order == "desc"

	sort.Slice(chars, func(i, j int) bool {
		a, b := chars[i], chars[j]

		var less bool
		switch sortBy {
		case "ki":
			less = kiToInt(a.Ki) < kiToInt(b.Ki)
		case "maxki": // on normalise maxKi -> maxki
			less = kiToInt(a.MaxKi) < kiToInt(b.MaxKi)
		default: // "name"
			less = strings.ToLower(a.Name) < strings.ToLower(b.Name)
		}

		if desc {
			return !less
		}
		return less
	})
}

// /search : recherche par nom + filtres + tri + pagination + favoris
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	theme := r.URL.Query().Get("theme")
	themeClass := "theme-classic"
	themeParam := ""
	if theme == "ui" {
		themeClass = "theme-ui"
		themeParam = "theme=ui"
	}

	// Filtres
	name := r.URL.Query().Get("name")
	race := r.URL.Query().Get("race")
	affiliation := r.URL.Query().Get("affiliation")

	// Tri
	sortBy := r.URL.Query().Get("sort")
	if sortBy == "" {
		sortBy = "name"
	}
	order := r.URL.Query().Get("order")
	if order == "" {
		order = "asc"
	}

	// Pagination
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize <= 0 {
		pageSize = 24
	}

	// Résultats (via cache)
	// (ta fonction filtre déjà name/race/affiliation)
	results := fonction.ApiSearchCharacters(name, race, affiliation)

	// Tri AVANT pagination
	normalizedSort := strings.ToLower(strings.TrimSpace(sortBy))
	if normalizedSort == "maxki" || normalizedSort == "maxKi" {
		normalizedSort = "maxki"
	} else if normalizedSort != "ki" {
		normalizedSort = "name"
	}
	sortCharacters(results, normalizedSort, order)

	// Pagination (après tri)
	total := len(results)
	totalPages := (total + pageSize - 1) / pageSize
	if totalPages < 1 {
		totalPages = 1
	}
	if page > totalPages {
		page = totalPages
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pageItems := results[start:end]

	// Prev/Next (pour les liens)
	hasPrev := page > 1
	hasNext := page < totalPages
	prevPage := page - 1
	nextPage := page + 1
	if !hasPrev {
		prevPage = 1
	}
	if !hasNext {
		nextPage = totalPages
	}

	// Favoris
	favMap := favoritesToMap(loadFavorites())
	if favMap == nil {
		favMap = map[int]bool{}
	}

	data := struct_.SearchResultsData{
		Query:       strings.TrimSpace(name),
		Race:        race,
		Affiliation: affiliation,

		ThemeClass: themeClass,
		ThemeParam: themeParam,

		Sort:  sortBy,
		Order: order,

		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,

		HasPrev:  hasPrev,
		HasNext:  hasNext,
		PrevPage: prevPage,
		NextPage: nextPage,

		Favorites: favMap,
		Results:   pageItems,
	}

	renderTemplate(w, "RealSearch.html", data)
}
