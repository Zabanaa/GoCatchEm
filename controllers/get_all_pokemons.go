package controllers

import (
	"database/sql"
	"net/http"
	"pokemon_api/models"
)

func GetAllPokemons(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var response Response

	rows, err := db.Query("SELECT * FROM pokemons")

	if err != nil {

		response.Meta.Type = "internal server error"
		response.Meta.StatusCode = http.StatusInternalServerError
		response.Body = map[string]interface{}{"error": err.Error()}

		respondWithJson(w, http.StatusInternalServerError, response)
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

			response.Meta.Type = "internal server error"
			response.Meta.StatusCode = http.StatusInternalServerError
			response.Body = map[string]interface{}{"error": err.Error()}

			respondWithJson(w, http.StatusInternalServerError, response)
		}

		pokemons = append(pokemons, pokemon)
	}

	response.Meta.Type = "success"
	response.Meta.StatusCode = http.StatusOK
	response.Body = make(map[string]interface{})
	response.Body["count"] = len(pokemons)
	response.Body["data"] = pokemons

	// Respond with the pokemons
	respondWithJson(w, http.StatusOK, response)
}
