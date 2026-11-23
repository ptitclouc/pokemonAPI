package controllers

import (
	"API/helpers"
	"API/models"
	"API/services"
	"errors"
	"log"
	"net/http"
	"strings"
)

func ListePoke(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.URL.Query().Get("nom"))
	characters, err := services.SearchPokemon(search)
	if err != nil {
		renderAPIError(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "list", models.ListPage{Search: search, Characters: characters})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

func renderAPIError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound,
			"Ce personnage n'existe pas.")
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest,
			"Identifiant de personnage invalide.")
	case errors.Is(err, models.ErrUnavailable):
		log.Println("Pokémon API :", err)
		helpers.RenderError(w, http.StatusServiceUnavailable,
			"Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError,
			"Une erreur est survenue.")
	}
}

func ShowPokemon(w http.ResponseWriter, r *http.Request) {
	character, err := services.GetPokemon(r.PathValue("id"))
	if err != nil {
		renderAPIError(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "character", character)
}
