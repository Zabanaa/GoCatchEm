package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func DeletePokemon(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pokemonName := vars["name"]
	var response Response

	query := `DELETE FROM pokemons WHERE name = $1;`
	result, err := db.Exec(query, pokemonName)

	if err != nil {

		response.ServerError(w, err.Error())
		return
	}

	count, err := result.RowsAffected()

	if err != nil {
		response.NotFound(w, "Pokemon not found.")
		return
	}

	if count == 0 {
		response.NotFound(w, "Pokemon not found.")
		return
	}

	response.Deleted(w)

}
