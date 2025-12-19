package fonction

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
    "Livrable-projet-groupie-tracker/struct"
	
)


var Data interface{}
var DecodeChar struct_.Characters
var DecodePlan struct_.Planets


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
