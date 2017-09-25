package controllers

import (
	"database/sql"
	"net/http"
	"pokemon_api/models"

	"github.com/gorilla/mux"
)

func GetPokemonsByType(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var response Response
	var errorMessage string
	var query string

	vars := mux.Vars(r)
	pokemonType := vars["type"]

	query = "SELECT * FROM pokemons WHERE types LIKE '%' || $1 || '%' ORDER BY id;"

	rows, err := db.Query(query, pokemonType)

	if err != nil {

		errorMessage = err.Error()
		response.ServerError(w, errorMessage)
	}

	defer rows.Close()

	var pokemons []models.Pokemon

	for rows.Next() {

		var pokemon models.Pokemon

		err := rows.Scan(
			&pokemon.ID, &pokemon.Number, &pokemon.Name, &pokemon.JpName,
			&pokemon.Types, &pokemon.Stats.HP, &pokemon.Stats.Attack,
			&pokemon.Stats.Defense, &pokemon.Stats.Sp_atk, &pokemon.Stats.Sp_def,
			&pokemon.Stats.Speed, &pokemon.Bio, &pokemon.Generation)

		if err != nil {

			errorMessage = err.Error()
			response.ServerError(w, errorMessage)
		}

		pokemons = append(pokemons, pokemon)
	}

	if len(pokemons) == 0 {
		response.NotFound(w, "This type does not exist")
		return
	}

	body := map[string]interface{}{"count": len(pokemons), "data": pokemons}
	response.StatusOK(w, body)

}
