package routes

import (
	"API/api"
	"encoding/json"
	"net/http"
)
 
func route() {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/pokemon", http.StatusMovedPermanently)
    })
    mux.HandleFunc("GET /pokemon", func(w http.ResponseWriter, r *http.Request) {
        pokemonList, err := api.GetAllPokemon() 
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(pokemonList)
    })
}
 