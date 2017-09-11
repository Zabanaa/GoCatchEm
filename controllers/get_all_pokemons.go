package controllers

import (
	"database/sql"
	"net/http"
	"pokemon_api/models"
)

func GetAllPokemons(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	pokemons, err := models.GetPokemons(db)

	if err != nil {
		// repond with error
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	// Respond with the pokemons
	respondWithJson(w, http.StatusOK, pokemons)
}
