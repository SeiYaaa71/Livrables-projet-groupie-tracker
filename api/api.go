package Api

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type ApiData struct {
    Results []interface{} `json:"data"`
}

func main() {
    // URL de l'API
    urlApi := "https://api.pokemontcg.io/v2/cards" // L'API nécessite un endpoint valide

    // Création de la requête HTTP
    req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
    if errReq != nil {
        fmt.Println("Oupss, une erreur est survenue :", errReq.Error())
        return
    }

    // Ajout d’un User-Agent
    req.Header.Add("X-Api-Key", "d2804cc0-062f-47ad-9b3d-41780ba2b8f6")

    // Exécution de la requête HTTP
    client := &http.Client{} // <-- correction ici
    res, errResp := client.Do(req)
    if errResp != nil {
        fmt.Println("Oupss, une erreur est survenue :", errResp.Error())
        return
    }
    defer res.Body.Close()

    // Lecture du corps de la réponse
    body, errBody := io.ReadAll(res.Body)
    if errBody != nil {
        fmt.Println("Oupss, une erreur est survenue :", errBody.Error())
        return
    }

    // Variable qui va contenir les données
    var decodeData ApiData

	fmt.Println("Réponse brute :")
	fmt.Println(string(body))

    // Décodage du JSON
    errJson := json.Unmarshal(body, &decodeData)
    if errJson != nil {
        fmt.Println("Erreur lors du décodage JSON :", errJson.Error())
        return
    }

    // Affichage des données
    if len(decodeData.Results) > 0 {
        fmt.Println(decodeData.Results[0])
    } else {
        fmt.Println("Aucun résultat trouvé")
    }
}