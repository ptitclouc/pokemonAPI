package api

import (
	"API/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// BaseURL est l'URL de base de l'API Tyradex pour les Pokémon.
const BaseURL = "https://tyradex.app/api/v1"

var client = &http.Client{
	Timeout: 5 * time.Second, // délai maximum pour la requête complète
}

func GetPokemon(id int) (models.Pokemon, error) {
	endpoint := fmt.Sprintf("%s/pokemon/%d", BaseURL, id)
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return models.Pokemon{}, err
	}
	request.Header.Set("User-Agent", "DataDex/1.0")

	response, err := client.Do(request)
	if err != nil {
		return models.Pokemon{}, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return models.Pokemon{}, models.ErrNotFound
	}

	if response.StatusCode != http.StatusOK {
		return models.Pokemon{}, fmt.Errorf("réponse inattendue : %s",
			response.Status)
	}

	var pokemon models.Pokemon
	err = json.NewDecoder(response.Body).Decode(&pokemon)
	return pokemon, err
}

func GetAllPokemon(name string) ([]models.Pokemon, error) {
	endpoint := fmt.Sprintf("%s/pokemon", BaseURL)
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "DataDex/1.0")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return []models.Pokemon{}, models.ErrNotFound
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("réponse inattendue : %s",
			response.Status)
	}
	var pokemonList []models.Pokemon
	err = json.NewDecoder(response.Body).Decode(&pokemonList)
	return pokemonList, err

}


