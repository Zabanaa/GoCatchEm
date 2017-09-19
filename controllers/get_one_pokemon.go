package controllers

import (
	"database/sql"
	"net/http"
	"pokemon_api/models"

	"github.com/gorilla/mux"
)

func GetPokemon(db *sql.DB, w http.ResponseWriter, req *http.Request) {

	var response Response
	var errorMessage string

	query := "SELECT * FROM pokemons WHERE name=$1 ORDER BY id"
	vars := mux.Vars(req)
	pokemonName := vars["name"]

	var pokemon models.Pokemon

	err := db.QueryRow(query, pokemonName).Scan(
		&pokemon.ID, &pokemon.Number, &pokemon.Name, &pokemon.JpName,
		&pokemon.Types, &pokemon.Stats.HP, &pokemon.Stats.Attack,
		&pokemon.Stats.Defense, &pokemon.Stats.Sp_atk, &pokemon.Stats.Sp_def,
		&pokemon.Stats.Speed, &pokemon.Bio, &pokemon.Generation)

	if err != nil {

		if err == sql.ErrNoRows {
			response.NotFound(w)
			return
		} else {
			errorMessage = err.Error()
			response.ServerError(w, errorMessage)
			return
		}
	}

	body := map[string]interface{}{"pokemon": pokemon}
	response.StatusOK(w, body)
}
