# 🐉 Dragon Ball Characters Explorer --- Go Web App

Cette application web, développée en **Go**, permet d'explorer les
personnages de l'univers **Dragon Ball** grâce à l'API officielle :\
👉 https://web.dragonball-api.com/

Elle intègre :

-   🔍 **Recherche** de personnages\
-   🧪 **Filtres avancés** (race, genre, affiliation, etc.)\
-   📄 **Pagination** par groupes de 10 ressources\
-   ⭐ **Système de favoris persistant**\
-   🔐 **Authentification OAuth2** via token d'accès\
-   🌐 Interface HTML simple (templates)

## 📁 Structure du projet

    .
    ├── api
    │   └── api.go                # Gestion du token et des appels API
    ├── controller
    │   └── controller.go         # Logique principale (recherche, filtres, pagination)
    ├── router
    │   └── router.go             # Définition des routes
    ├── templetes
    │   └── index.html            # Vue HTML
    ├── main.go                   # Lancement du serveur
    └── go.mod

## 🔧 Installation et configuration

### 1. Cloner le repository

``` sh
git clone <url-du-projet>
cd dragonball-app
```

## 🔐 Configuration des variables d'environnement

``` sh
export API_CLIENT_ID="TON_CLIENT_ID"
export API_CLIENT_SECRET="TON_CLIENT_SECRET"
```

Recharge l'environnement :

``` sh
source ~/.bashrc
# ou
source ~/.zshrc
```

## 🚀 Lancer l'application

``` sh
go mod tidy
go run main.go
```

Disponible sur **http://localhost:8080**

## 🧠 Fonctionnalités

-   Recherche de personnages\
-   Filtres avancés : race, genre, affiliation\
-   Pagination par 10\
-   Favoris persistants (`favorites.json`)

## 📦 Appels API

`api.go` gère : - Token OAuth2\
- Stockage temporaire\
- Requêtes API Dragon Ball

## 🛠️ Technologies

  Technologie       Usage
  ----------------- ----------
  Go                Backend
  net/http          Serveur
  html/template     Vue HTML
  OAuth2            Auth
  JSON              Favoris
  Dragon Ball API   Données

## 📚 Améliorations futures

-   Cache API\
-   Fiche personnage détaillée\
-   UI améliorée\
-   Sessions utilisateur

## 📄 Licence

Projet libre d'utilisation pédagogique.
