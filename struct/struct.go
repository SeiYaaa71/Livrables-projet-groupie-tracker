package struct_

type CharactersByName struct {
}

// Character définit un personnage unique avec les champs de ton API

type CharacterById struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Ki          string  `json:"ki"`
	MaxKi       string  `json:"maxKi"`
	Race        string  `json:"race"`
	Gender      string  `json:"gender"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Affiliation string  `json:"affiliation"`
	DeletedAt   *string `json:"deletedAt"`

	OriginPlanet struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		IsDestroyed bool    `json:"isDestroyed"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		DeletedAt   *string `json:"deletedAt"`
	} `json:"originPlanet"`

	Transformations []struct {
		ID        int     `json:"id"`
		Name      string  `json:"name"`
		Image     string  `json:"image"`
		Ki        string  `json:"ki"`
		DeletedAt *string `json:"deletedAt"`
	} `json:"transformations"`
}

// Characters est la structure de retour de l'API pour une liste
type Characters struct {
	Items []CharacterById `json:"items"`
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
	Query      string
	Results    []CharacterById
}

type SearchResultsData struct {
  Query       string
  Race        string
  Affiliation string

  ThemeClass string
  ThemeParam string

  Sort  string
  Order string

  Page       int
  PageSize   int
  Total      int
  TotalPages int

  HasPrev  bool
  HasNext  bool
  PrevPage int
  NextPage int

  Favorites map[int]bool
  Results   []CharacterById
}

type FavoritesPageData struct {
	ThemeClass string
	ThemeParam string

	Favorites map[int]bool
	Results   []CharacterById
	Total     int
}


var Filters = [][]string{
	{"race", "Human", "Saiyan", "Namekian", "Majin", "Frieza Race", "Android", "Jiren Race", "God", "Angel", "Evil", "Nucleico", "Nucleico benigno", "Unknown"},
	{"affiliation", "Z Fighter", "Red Ribbon Army", "Namekian Warrior", "Freelancer", "Army of Frieza", "Pride Troopers", "Assistant of Vermoud", "God", "Assistant of Beerus", "Villain", "Other"},
}
