package services

import (
	"API/api"
	"API/models"
	"errors"
	"strconv"
)

// SearchCharacters renvoie les personnages de la première page,
// filtrés par l'API sur le nom si une recherche est faite.
func SearchPokemon(name string) ([]models.Pokemon, error) {
	page, err := api.GetPokemon(name)
	// Une recherche sans résultat n'est pas une erreur pour notre site :
	// l'API répond 404, on renvoie simplement une liste vide.
	if errors.Is(err, models.ErrNotFound) && name != "" {
		return []models.Pokemon{}, nil
	}
	if err != nil {
		return nil, err
	}
	return page.Results, nil
}

// GetCharacter vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetPokemon(idParam string) (models.Pokemon, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return models.Pokemon{}, models.ErrBadRequest
	}
	return api.GetPokemon(id)
}
