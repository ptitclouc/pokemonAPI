package models

type ListPage struct {
	Search     string    // recherche saisie, pour la réafficher dans le formulaire
	Characters []Pokemon `json:"result"`
}

type Pokemon struct {
	Id       int     `json:"pokedex_id"`
	Category string  `json:"category"`
	Name     Name    `json:"name"`
	Sprites  Sprites `json:"sprites"`
	Types    []Type  `json:"types"`
}

type Name struct {
	Fr string `json:"fr"`
}

type Sprites struct {
	Regular string `json:"regular"`
}

type Type struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
