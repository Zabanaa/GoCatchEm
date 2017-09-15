package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DeletePokemon(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pokemonName := vars["name"]
	var response Response

	query := `DELETE FROM pokemons WHERE id=$1`
	result, err := db.Exec(query, pokemonName)

	if err != nil {

		if err == sql.ErrNoRows {
			response.NotFound(w, "Pokemon not found.")
			return
		} else {
			response.ServerError(w, err.Error())
			return
		}
	}

	count, err := result.RowsAffected()

	if err != nil {
		// return an error ?
		fmt.Println(err.Error())
	}

	log.Println("Deleted", count, "rows")

	response.Deleted(w)

}
