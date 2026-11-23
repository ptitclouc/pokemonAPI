package services

import (
	"API/api"
	"API/models"
	"errors"
	"strconv"
	"strings"
)

// SearchCharacters renvoie les personnages de la première page,
// filtrés par l'API sur le nom si une recherche est faite.
func SearchPokemon(name string) ([]models.Pokemon, error) {
	page, err := api.GetAllPokemon(name)
	if errors.Is(err, models.ErrNotFound) && name != "" {
		return []models.Pokemon{}, nil
	}
	if err != nil {
		return nil, err
	}
	if name == "" {
		return page, nil
	}

	name = strings.ToLower(name)
	filtered := []models.Pokemon{}
	for _, p := range page {
		if p.Id != 0 && strings.Contains(strings.ToLower(p.Name.Fr), name) {
			filtered = append(filtered, p)
		}
	}
	return filtered, nil
}

// GetCharacter vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetPokemon(idParam string) (models.Pokemon, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return models.Pokemon{}, models.ErrBadRequest
	}
	return api.GetPokemon(id)
}
