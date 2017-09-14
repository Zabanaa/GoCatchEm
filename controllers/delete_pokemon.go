package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func deletePokemon(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	// Fetch a pokemon by ID
	// Delete the motherfucker
	// return a 204
	fmt.Fprintf(w, "hello world")
}
