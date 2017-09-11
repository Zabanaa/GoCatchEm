package controllers

import (
	"fmt"
	"net/http"
)

func updatePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
