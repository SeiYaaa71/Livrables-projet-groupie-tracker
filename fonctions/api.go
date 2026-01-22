package fonction

import (
	struct_ "Livrable-projet-groupie-tracker/struct"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var Data interface{}

var nonDigits = regexp.MustCompile(`[^0-9]`)

var (
	charactersCache []struct_.CharacterById
	cacheMu         sync.RWMutex
)

func ApiGet() ([]struct_.CharacterById, error) {
	var all []struct_.CharacterById
	page := 1
	limit := 50

	for {
		url := fmt.Sprintf("characters?page=%d&limit=%d", page, limit)

		baseURL := "https://dragonball-api.com/api/"
		fullURL := baseURL + url

		resp, err := http.Get(fullURL)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}

		// La réponse attendue ressemble à: { "items": [...] }
		var parsed struct_.Characters
		if err := json.Unmarshal(body, &parsed); err != nil {
			return nil, err
		}

		if len(parsed.Items) == 0 {
			break
		}

		all = append(all, parsed.Items...)
		page++
	}

	return all, nil
}

func ApiSearchCharacters(name, race, affiliation string) []struct_.CharacterById {
	cacheMu.RLock()
	all := charactersCache
	cacheMu.RUnlock()

	name = strings.ToLower(strings.TrimSpace(name))

	var results []struct_.CharacterById
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

func LoadCharactersCache() error {
	all, err := ApiGet()
	if err != nil {
		return err
	}

	cacheMu.Lock()
	charactersCache = all
	cacheMu.Unlock()

	fmt.Println("Cache chargé :", len(all), "personnages")
	return nil
}

// Retourne une COPIE du cache (pratique pour éviter qu’on modifie la slice globale)
func GetCharactersCached() []struct_.CharacterById {
	cacheMu.RLock()
	defer cacheMu.RUnlock()

	out := make([]struct_.CharacterById, len(charactersCache))
	copy(out, charactersCache)
	return out
}

func GetCharacterByID(id int) (struct_.CharacterById, bool) {
	if id < 1 {
		return struct_.CharacterById{}, false
	}

	url := fmt.Sprintf("https://dragonball-api.com/api/characters/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return struct_.CharacterById{}, false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return struct_.CharacterById{}, false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return struct_.CharacterById{}, false
	}

	var c struct_.CharacterById
	if err := json.Unmarshal(body, &c); err != nil {
		return struct_.CharacterById{}, false
	}

	// sécurité: si l’API renvoie un objet sans ID
	if c.ID < 1 {
		return struct_.CharacterById{}, false
	}

	return c, true
}

func kiToInt(s string) int {
	// "60.000.000" / "unknown" -> int
	cleaned := nonDigits.ReplaceAllString(s, "")
	if cleaned == "" {
		return 0
	}
	n, err := strconv.Atoi(cleaned)
	if err != nil {
		return 0
	}
	return n
}

func SortCharacters(chars []struct_.CharacterById, sortBy, order string) {
	desc := (order == "desc")

	sort.Slice(chars, func(i, j int) bool {
		a, b := chars[i], chars[j]

		var less bool
		switch sortBy {
		case "ki":
			less = kiToInt(a.Ki) < kiToInt(b.Ki)
		case "maxKi":
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
