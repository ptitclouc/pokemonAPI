# Pokédex

Petit site web développé en **Go** qui permet de consulter et de rechercher des Pokémon.

Les données sont récupérées depuis l’API [Tyradex](https://tyradex.app/).

## Fonctionnalités

- Afficher la liste des Pokémon.
- Rechercher un Pokémon par son nom.
- Consulter la fiche d’un Pokémon : image, numéro, catégorie et types.
- Afficher des pages d’erreur en cas de problème.

## Prérequis

- Go (version indiquée dans `go.mod`).

## Lancer le projet

1. Ouvrir un terminal dans le dossier `pokemonAPI`.
2. Lancer le serveur :

   ```bash
   go run .
   ```

3. Ouvrir cette adresse dans un navigateur :

   http://localhost:8080

## Routes principales

| Adresse | Description |
|---|---|
| `/personnages` | Affiche la liste des Pokémon |
| `/personnages?nom=pikachu` | Recherche un Pokémon par son nom |
| `/personnages/25` | Affiche la fiche du Pokémon portant l’identifiant 25 |

## Structure du projet

- `api/` : appels à l’API Tyradex.
- `controllers/` : gestion des requêtes HTTP.
- `helpers/` : rendu des templates HTML.
- `models/` : structures de données.
- `routes/` : définition des routes.
- `services/` : logique de recherche et de récupération des Pokémon.
- `templates/` : pages HTML.
- `assets/` : fichiers statiques, dont le CSS.

## Arrêter le serveur

Dans le terminal, appuyer sur `Ctrl + C`.
