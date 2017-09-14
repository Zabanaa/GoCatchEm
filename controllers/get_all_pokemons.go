package controllers

import (
	"database/sql"
	"net/http"
	"pokemon_api/models"
)

func GetAllPokemons(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var response Response
	var errorMessage string

	rows, err := db.Query("SELECT * FROM pokemons")

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

	body := map[string]interface{}{"count": len(pokemons), "data": pokemons}
	response.StatusOK(w, body)

}
