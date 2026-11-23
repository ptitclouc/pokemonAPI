package helpers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

// LoadTemplates charge tous les templates une seule fois, au démarrage.
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

// Render exécute le template dans un buffer, puis envoie la page et son code.
func Render(w http.ResponseWriter, status int, name string, data any) {
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, name, data); err != nil {
		log.Println("rendu du template :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}

// ErrorPage : données affichées par le template d'erreur.
type ErrorPage struct {
	Code    int
	Message string
}

// RenderError affiche la page d'erreur commune avec le bon code HTTP.
func RenderError(w http.ResponseWriter, status int, message string) {
	Render(w, status, "error", ErrorPage{Code: status, Message: message})
}
