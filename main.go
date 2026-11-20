package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

// getPokemon interroge l'API pour récupérer un personnage par son id.
// Chaque échec retourne directement l'erreur personnalisée correspondante.
func getPokemon(client *http.Client, id string) (*Pokemon, error) {
	endpoint := fmt.Sprintf("%s/%v", baseURL, id)

	// Préparation de la requête GET.
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, ErrRequestCreation
	}

	// Envoi de la requête.
	response, err := client.Do(request)
	if err != nil {
		return nil, ErrRequestSend
	}
	defer response.Body.Close()

	// Vérification du code de statut HTTP.
	switch response.StatusCode {
	case http.StatusOK:
		// Succès : on continue vers le décodage.
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusInternalServerError:
		return nil, ErrInvalidID
	default:
		return nil, ErrUnexpectedState
	}

	// Décodage du JSON dans la structure Pokemon.
	var Pokemon Pokemon
	if err := json.NewDecoder(response.Body).Decode(&Pokemon); err != nil {
		return nil, ErrDecode
	}

	return &Pokemon, nil
}

func main() {
	// Client HTTP avec un timeout pour éviter de bloquer indéfiniment.
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	id := "fdfdvfdf"

	Pokemon, err := getPokemon(client, id)
	if err != nil {
		// Gestion différente selon l'erreur retournée.
		switch {
		case errors.Is(err, ErrNotFound):
			log.Fatalf("Personnage %q introuvable, vérifiez l'id.", id)
		case errors.Is(err, ErrInvalidID):
			log.Fatalf("L'id %q n'est pas valide (un nombre est attendu).", id)
		case errors.Is(err, ErrRequestCreation):
			log.Fatalf("Problème interne lors de la création de la requête.")
		case errors.Is(err, ErrRequestSend):
			log.Fatalf("Impossible de joindre l'API, vérifiez votre connexion.")
		case errors.Is(err, ErrDecode):
			log.Fatalf("La réponse de l'API est illisible.")
		case errors.Is(err, ErrUnexpectedState):
			log.Fatalf("L'API a renvoyé une réponse inattendue.")
		default:
			log.Fatalf("Erreur inconnue : %v", err)
		}
	}

	fmt.Println(*Pokemon)
}
