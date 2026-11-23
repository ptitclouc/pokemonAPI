package routes

import (
	"API/controllers"
	"net/http"
)

func MainRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fichiers statiques : /static/css/style.css → assets/css/style.css
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// L'accueil redirige vers la liste des personnages (redirection permanente)
	mux.Handle("GET /{$}", http.RedirectHandler("/pokemon", http.StatusMovedPermanently))

	PokemonRoutes(mux)

	// Toute autre URL : page 404 du site
	mux.HandleFunc("/", controllers.NotFound)
	return mux
}
