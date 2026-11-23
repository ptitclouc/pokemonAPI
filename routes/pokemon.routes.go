package routes

import "net/http"

func characterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /personnages", controllers.ListPokemon)
	mux.HandleFunc("GET /personnages/{id}", controllers.ShowCharacter)
}
