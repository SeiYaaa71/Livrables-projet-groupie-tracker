package api

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var baseURL = "https://web.dragonball-api.com"

// Récupération du Client ID
func getClientID() string {
	id := os.Getenv("API_CLIENT_ID")
	if id == "" {
		fmt.Println("⚠️  API_CLIENT_ID est vide !")
	}
	return id
}

// Récupération du Client Secret
func getClientSecret() string {
	secret := os.Getenv("API_CLIENT_SECRET")
	if secret == "" {
		fmt.Println("⚠️  API_CLIENT_SECRET est vide !")
	}
	return secret
}

// Génération du header "Authorization: Basic "
func getBasicAuthHeader() string {
	id := getClientID()
	secret := getClientSecret()

	token := id + ":" + secret
	encoded := base64.StdEncoding.EncodeToString([]byte(token))

	return "Basic " + encoded
}

// Fonction générique d'appel API
func CallAPI(endpoint string) ([]byte, error) {
	url := baseURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Ajout de l'authentification Basic (Client ID + Secret)
	req.Header.Set("Authorization", getBasicAuthHeader())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
