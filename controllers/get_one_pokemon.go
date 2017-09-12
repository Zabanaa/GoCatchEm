package controllers

import (
	"database/sql"
	"net/http"
	"pokemon_api/models"

	"github.com/gorilla/mux"
)

func GetPokemon(db *sql.DB, w http.ResponseWriter, req *http.Request) {

	var response Response

	vars := mux.Vars(req)
	pokemonName := vars["name"]

	var pokemon models.Pokemon

	err := db.QueryRow("SELECT * FROM pokemons WHERE name=$1", pokemonName).Scan(
		&pokemon.ID, &pokemon.Number, &pokemon.Name, &pokemon.JpName,
		&pokemon.Types, &pokemon.Stats.HP, &pokemon.Stats.Attack,
		&pokemon.Stats.Defense, &pokemon.Stats.Sp_atk, &pokemon.Stats.Sp_def,
		&pokemon.Stats.Speed, &pokemon.Bio, &pokemon.Generation)

	if err != nil {

		if err == sql.ErrNoRows {
			// 404 not found
			// respondWithError(w, http.StatusNotFound, "Pokemon not found")
			return
		} else {
			// 500 Internal Server Error
			// respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	response.Body = make(map[string]interface{})
	response.Body["pokemon"] = pokemon
	respondWithJson(w, http.StatusOK, response)
}
