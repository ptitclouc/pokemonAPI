package routes

import (
	"API/controllers"
	"net/http"
)

func PokemonRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /pokemon", controllers.ListePoke)
	mux.HandleFunc("GET /pokemon/{id}", controllers.ShowPokemon)
}
