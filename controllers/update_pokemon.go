package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func UpdatePokemon(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	// Get the pokemon by name
	// fetch it from the db
	// update the motherfucker
	// return the new object
	fmt.Fprintf(w, "hello world")
}
