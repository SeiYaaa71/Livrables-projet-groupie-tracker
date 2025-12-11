package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// Structure correspondant à la réponse JSON de l'API
type ApiData struct {
    Items []interface{} `json:"items"` // L’API renvoie "items", pas "data"
}

func main() {

    // ⚠️ L’endpoint racine ne renvoie pas de données utiles.
    // Utilise plutôt /api/characters pour obtenir des résultats.
    urlApi := "https://dragonball-api.com/api/characters"

    // Création de la requête
    req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
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
    var decodeData ApiData

    errJson := json.Unmarshal(body, &decodeData)
    if errJson != nil {
        fmt.Println("Erreur lors du décodage JSON :", errJson.Error())
        return
    }

    // Affichage du premier élément
    if len(decodeData.Items) > 0 {
        fmt.Println("Premier élément :", decodeData.Items[0])
    } else {
        fmt.Println("Aucun résultat trouvé")
    }
}
