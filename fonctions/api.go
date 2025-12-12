package fonction

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
	
)

// Structures correspondant aux réponses JSON de l'API
type characters struct {
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

type planets struct {
	Items []struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		IsDestroyed bool    `json:"isDestroyed"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		DeletedAt   *string `json:"deletedAt"`
	} `json:"items"`
}


var Data interface{}
var DecodeChar characters
var DecodePlan planets


var Filters = [][]string{
    //races filters
    {"race", "Human", "Saiyan", "Namekian", "Majin", "Frieza Race", "Android", "Jiren Race", "God", "Angel", "Evil", "Nucleico", "Nucleico benigno", "Unknown"},

    //affiliation
    {"affiliation", "Z Fighter", "Red Ribbon Army", "Namekian Warrior", "Freelancer", "Army of Frieza", "Pride Troopers", "Assistant of Vermoud", "God", "Assistant of Beerus", "Villain", "Other"},

}

func ApiGet(Url string, filters []string) {

    UrlApi := "https://dragonball-api.com/api/"

    if len(filters) > 0 {
        UrlApi += "?" + strings.Join(filters, "&")
    } 


    // Création de la requête
    req, errReq := http.NewRequest(http.MethodGet, UrlApi, nil)
    if errReq != nil {
        fmt.Println("Erreur création requête :", errReq.Error())
        return
    }

    client := &http.Client{}
    res, errResp := client.Do(req)
    if errResp != nil {
        fmt.Println("Erreur lors de l'appel API :", errResp.Error())
        return
    }
    defer res.Body.Close()

    // Lecture du corps de la réponse
    body, errBody := io.ReadAll(res.Body)
    if errBody != nil {
        fmt.Println("Erreur lecture body :", errBody.Error())
        return
    }

    fmt.Println("Réponse brute :")
    fmt.Println(string(body))

    // Décodage JSON
   

    if Url == "characters"{
        errJson := json.Unmarshal(body, &DecodeChar)
        if errJson != nil {
            fmt.Println("Erreur lors du décodage JSON :", errJson.Error())
            fmt.Println("Contenu reçu :", string(body))
            return
        }

        Data = &DecodeChar
    } else {
        errJson := json.Unmarshal(body, &DecodePlan)
        if errJson != nil {
            fmt.Println("Erreur lors du décodage JSON :", errJson.Error())
            fmt.Println("Contenu reçu :", string(body))
            return
        }

        Data = &DecodePlan        
    }
}
