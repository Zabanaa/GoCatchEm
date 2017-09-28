package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pokemon_api/app/models"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
)

func UpdatePokemon(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var response Response
	var errorMessage string
	var pokemon models.Pokemon
	var requestBody models.Pokemon

	// 1. Get pokemon by name from the DB
	query := "SELECT * FROM pokemons WHERE name = $1"
	vars := mux.Vars(r)
	pokemonName := vars["name"]

	err := db.QueryRow(query, pokemonName).Scan(&pokemon.ID, &pokemon.Number,
		&pokemon.Name, &pokemon.JpName, &pokemon.Types, &pokemon.Stats.HP,
		&pokemon.Stats.Attack, &pokemon.Stats.Defense, &pokemon.Stats.Sp_atk,
		&pokemon.Stats.Sp_def, &pokemon.Stats.Speed, &pokemon.Bio, &pokemon.Generation)

	if err != nil {

		if err == sql.ErrNoRows {
			response.NotFound(w, "This Pokemon does not exist")
			return
		} else {
			errorMessage = err.Error()
			response.ServerError(w, errorMessage)
			return
		}
	}

	// 2. Store the payload in another pokemon struct
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	if err := decoder.Decode(&requestBody); err != nil {
		response.BadRequest(w)
		return
	}

	// 3. loop through the body struct fields
	pokemonFromRequest := reflect.ValueOf(&requestBody).Elem()
	pokemonToBeUpdated := reflect.ValueOf(&pokemon).Elem()

	for i := 1; i < pokemonFromRequest.NumField(); i++ {

		pokemonField := pokemonToBeUpdated.Field(i)
		requestBodyField := pokemonFromRequest.Field(i)

		switch requestBodyField.Kind() {

		case reflect.Ptr:
			if !reflect.Indirect(requestBodyField).IsValid() {
				// Value is not in the json body
				// Skip to the next element
				continue
			}
		case reflect.Struct:
			// loop through the Stats fields and check the pointers
			continue
		}

		pokemonField.Set(reflect.Value(requestBodyField))

	}

	body := map[string]interface{}{"pokemon": pokemon}
	response.StatusOK(w, body)

	// exec update statement using the new fields
}

func ActuallyUpdatePokemon(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var response Response
	var errorMessage string
	var pokemon models.Pokemon
	var requestBody models.Pokemon

	// 1. Get pokemon by name from the DB
	query := "SELECT * FROM pokemons WHERE name = $1"
	vars := mux.Vars(r)
	pokemonName := vars["name"]

	err := db.QueryRow(query, pokemonName).Scan(&pokemon.ID, &pokemon.Number,
		&pokemon.Name, &pokemon.JpName, &pokemon.Types, &pokemon.Stats.HP,
		&pokemon.Stats.Attack, &pokemon.Stats.Defense, &pokemon.Stats.Sp_atk,
		&pokemon.Stats.Sp_def, &pokemon.Stats.Speed, &pokemon.Bio, &pokemon.Generation)

	if err != nil {

		if err == sql.ErrNoRows {
			response.NotFound(w, "This pokemon does not exist")
			return

		} else {
			errorMessage = err.Error()
			response.ServerError(w, errorMessage)
			return
		}
	}

	// 2. Store the payload in another pokemon struct
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	if err := decoder.Decode(&requestBody); err != nil {
		response.BadRequest(w)
		return
	}

	// 3. loop through the body struct fields
	pokemonFromRequest := reflect.ValueOf(&requestBody).Elem()
	pokemonToBeUpdated := reflect.ValueOf(&pokemon).Elem()

	for i := 1; i < pokemonFromRequest.NumField(); i++ {

		pokemonField := pokemonToBeUpdated.Field(i)
		requestBodyField := pokemonFromRequest.Field(i)

		switch requestBodyField.Kind() {

		case reflect.Ptr:
			if !reflect.Indirect(requestBodyField).IsValid() {
				// Value is not in the json body
				// Skip to the next element
				continue
			}
		case reflect.Struct:
			// loop through the Stats fields and check the pointers
			continue
		}

		pokemonField.Set(reflect.Value(requestBodyField))

	}

	// exec update statement using the new fields
	updateQuery := `
	UPDATE pokemons
	SET	number = $1, name = $2, jp_name = $3, types = $4, hp = $5, attack = $6,
	defense = $7, sp_atk = $8, sp_def = $9, speed = $10, bio = $11,
	generation = $12
	WHERE name = $13
	`

	_, err = db.Exec(updateQuery, pokemon.Number, pokemon.Name, pokemon.JpName,
		pokemon.Types, pokemon.Stats.HP, pokemon.Stats.Attack, pokemon.Stats.Defense,
		pokemon.Stats.Sp_atk, pokemon.Stats.Sp_def, pokemon.Stats.Speed, pokemon.Bio, pokemon.Generation, pokemonName)

	if err != nil {

		errorMessage := err.Error()

		if strings.Contains(errorMessage, "unique constraint") {

			response.Conflict(w)
			return

		} else {
			response.ServerError(w, errorMessage)
			return
		}
	}

	// return the object
	body := map[string]interface{}{"pokemon": pokemon}
	response.StatusOK(w, body)
}
