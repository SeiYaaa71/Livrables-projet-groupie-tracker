package controller

import "net/http"

// ResolveTheme lit ?theme=ui et retourne :
// - ThemeClass : "theme-ui" ou "theme-classic"
// - ThemeParam : "theme=ui" ou ""
//
// Utilisation : themeClass, themeParam := ResolveTheme(r)
func ResolveTheme(r *http.Request) (themeClass string, themeParam string) {
	theme := r.URL.Query().Get("theme")

	themeClass = "theme-classic"
	themeParam = ""

	if theme == "ui" {
		themeClass = "theme-ui"
		themeParam = "theme=ui"
	}

	return themeClass, themeParam
}
