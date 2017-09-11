package controllers

import (
	"fmt"
	"net/http"
)

func getPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
