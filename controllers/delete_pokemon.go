package controllers

import (
	"fmt"
	"net/http"
)

func deletePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
