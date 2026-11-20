package main

import (
	"API/api"
	"fmt"
)

func main() {
	fmt.Println(api.GetPokemon(25))
}