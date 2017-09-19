package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pokemon_api/models"
	"reflect"

	"github.com/gorilla/mux"
)

func UpdatePokemon(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var response Response
	var errorMessage string
	var pokemon models.Pokemon
	var updatedPokemon models.Pokemon

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
			response.NotFound(w)
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

	if err := decoder.Decode(&updatedPokemon); err != nil {
		response.BadRequest(w)
		return
	}

	// 3. loop through the body struct fields
	pokemonReflect := reflect.ValueOf(&pokemon).Elem()
	updatedPokemonReflect := reflect.ValueOf(&updatedPokemon).Elem()

	for i := 1; i < pokemonReflect.NumField(); i++ {

		pokemonReflectField := pokemonReflect.Field(i)
		updatedPokemonReflectField := updatedPokemonReflect.Field(i)

		// if nil, skip
		if updatedPokemonReflectField.Type() == nil {
			// continue
			pokemonReflectField.Set(reflect.Value(pokemonReflectField))
		} else {
			// else update the pokemon with the new field
			pokemonReflectField.Set(reflect.Value(updatedPokemonReflectField))
		}

	}

	// return the object
	body := map[string]interface{}{"pokemon": pokemon}
	response.StatusOK(w, body)

	// exec update statement using the new fields
}
