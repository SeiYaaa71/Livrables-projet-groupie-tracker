package fonction

import (
	struct_ "Livrable-projet-groupie-tracker/struct"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var Data interface{}

func ApiGet(url string, filters []string) {
	baseURL := "https://dragonball-api.com/api/"
	fullURL := baseURL + url

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Erreur HTTP :", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lecture body :", err)
		return
	}

	// Personnage par ID
	if strings.HasPrefix(url, "characters/") {
		var character struct_.CharacterById
		if err := json.Unmarshal(body, &character); err != nil {
			fmt.Println("Erreur JSON character :", err)
			return
		}
		Data = character
		return
	}

	// Liste personnages
	if url == "characters" {
		var characters struct_.Characters
		if err := json.Unmarshal(body, &characters); err != nil {
			fmt.Println("Erreur JSON characters :", err)
			return
		}
		Data = characters.Items
		return
	}
}

func ApiSearchCharacters(name, race, affiliation string) []struct_.CharacterById {

	ApiGet("characters", nil)

	all, ok := Data.([]struct_.CharacterById)
	if !ok {
		return nil
	}

	var results []struct_.CharacterById

	name = strings.ToLower(strings.TrimSpace(name))

	for _, c := range all {

		if name != "" && !strings.Contains(strings.ToLower(c.Name), name) {
			continue
		}

		if race != "" && c.Race != race {
			continue
		}

		if affiliation != "" && c.Affiliation != affiliation {
			continue
		}

		results = append(results, c)
	}

	return results
}
