package models

import "errors"

// Erreurs personnalisées, testables avec errors.Is.
var (
	ErrNotFound    = errors.New("ressource introuvable")
	ErrBadRequest  = errors.New("demande invalide")
	ErrUnavailable = errors.New("API indisponible")
)
