package struct_



type Characters struct {
	Items []struct {
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
	} `json:"items"`
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
    Results    []Characters
}

var Filters = [][]string{
    //races filters
    {"race", "Human", "Saiyan", "Namekian", "Majin", "Frieza Race", "Android", "Jiren Race", "God", "Angel", "Evil", "Nucleico", "Nucleico benigno", "Unknown"},

    //affiliation
    {"affiliation", "Z Fighter", "Red Ribbon Army", "Namekian Warrior", "Freelancer", "Army of Frieza", "Pride Troopers", "Assistant of Vermoud", "God", "Assistant of Beerus", "Villain", "Other"},

}