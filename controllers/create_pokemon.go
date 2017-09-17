package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pokemon_api/models"
	"strings"
)

func CreatePokemon(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var body models.Pokemon
	var response Response

	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	if err := decoder.Decode(&body); err != nil {
		response.BadRequest(w, `Couldn't process the request. Make sure it's properly formatted and that the fields are of the correct types. For more information on types, please refer to the documentation.`)
		return
	}

	query := `

	INSERT INTO pokemons (
		number, name, jp_name, types,
		hp, attack, defense, sp_atk, sp_def, speed,
		bio, generation)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);

	`
	_, err := db.Exec(query, &body.Number, &body.Name, &body.JpName,
		&body.Types, &body.Stats.HP, &body.Stats.Attack, &body.Stats.Defense,
		&body.Stats.Sp_atk, &body.Stats.Sp_def, &body.Stats.Speed, &body.Bio,
		&body.Generation)

	if err != nil {

		errorMessage := err.Error()

		if strings.Contains(errorMessage, "unique constraint") {

			response.Conflict(w)
			return

		} else if strings.Contains(errorMessage, "not-null") {

			missingField := extractMissingField(errorMessage)
			response.MissingFields(w, missingField)
			return

		} else {

			response.ServerError(w, errorMessage)
			return
		}
	}

	response.Created(w, "Pokemon Created")
}
