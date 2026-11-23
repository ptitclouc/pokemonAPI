package controllers

import (
	"API/models"
	"net/http"
	"strings"
	
)

type ListPage struct {
	Search     string // recherche saisie, pour la réafficher dans le formulaire
	Characters []models.Pokemon
}

func ListePoke(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.URL.Query().Get("nom"))
	characters, err := services.SearchCharacters(search)
	if err != nil {
		renderAPIError(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "list", ListPage{Search: search, Characters: characters})
}
