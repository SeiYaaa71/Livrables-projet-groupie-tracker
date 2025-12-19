package struct_

// Character définit un personnage unique avec les champs de ton API
type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Ki          string `json:"ki"`
	MaxKi       string `json:"maxKi"`
	Race        string `json:"race"`
	Gender      string `json:"gender"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Affiliation string `json:"affiliation"`
	DeletedAt   *string `json:"deletedAt"`
}

// Characters est la structure de retour de l'API pour une liste
type Characters struct {
	Items []Character `json:"items"`
}

type Planets struct {
	Items []struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		IsDestroyed bool    `json:"isDestroyed"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		DeletedAt   *string `json:"deletedAt"`
	} `json:"items"`
}

type SearchPageData struct {
	ThemeClass string
	ThemeParam string
	Results    []Character // Réfère maintenant à la structure définie plus haut
}

var Filters = [][]string{
	{"race", "Human", "Saiyan", "Namekian", "Majin", "Frieza Race", "Android", "Jiren Race", "God", "Angel", "Evil", "Nucleico", "Nucleico benigno", "Unknown"},
	{"affiliation", "Z Fighter", "Red Ribbon Army", "Namekian Warrior", "Freelancer", "Army of Frieza", "Pride Troopers", "Assistant of Vermoud", "God", "Assistant of Beerus", "Villain", "Other"},
}