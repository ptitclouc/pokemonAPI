package main

import (
	"API/api"
	"fmt"
	"log"
)

func main() {
	pokemon, err := api.GetPokemon(25)
	if err != nil {
		log.Fatalf("Impossible a charger le pokémon %v", err)
	}
	fmt.Println(pokemon)
}
