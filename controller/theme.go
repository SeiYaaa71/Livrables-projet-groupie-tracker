package controller

import "net/http"

// ResolveTheme lit ?theme=ui (ou ?theme=classic) et stocke le choix dans un cookie.
// Si aucun param n'est fourni, on relit le cookie pour garder le thème sur toutes les pages.
func ResolveTheme(r *http.Request) (themeClass string, themeParam string) {
	const cookieName = "theme"

	theme := r.URL.Query().Get("theme")

	// Si pas de theme dans l'URL, on essaie le cookie
	if theme == "" {
		if c, err := r.Cookie(cookieName); err == nil {
			theme = c.Value
		}
	}

	// Valeur par défaut
	themeClass = "theme-classic"
	themeParam = ""

	// Appliquer le thème
	if theme == "ui" {
		themeClass = "theme-ui"
		themeParam = "theme=ui"
	} else if theme == "classic" {
		themeClass = "theme-classic"
		themeParam = ""
	}

	// Si l'utilisateur force un thème via l'URL, on persiste dans un cookie
	// NOTE: on ne peut pas écrire le cookie ici (pas accès à ResponseWriter),
	// donc on laisse ResolveTheme "pur" et on écrit le cookie via middleware/handler.
	// Pour garder un changement minimal, on ajoute une fonction helper ci-dessous.
	return themeClass, themeParam
}

// PersistThemeCookie écrit le cookie theme=ui/classic (à appeler dans les handlers).
func PersistThemeCookie(w http.ResponseWriter, r *http.Request) {
	const cookieName = "theme"
	theme := r.URL.Query().Get("theme")
	if theme != "ui" && theme != "classic" {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    theme,
		Path:     "/",
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	})
}
