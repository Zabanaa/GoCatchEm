package main

import (
	"os"
)

func main() {

	var app App

	app.Initialize(
		os.Getenv("POKEDEX_DB_USER"),
		os.Getenv("POKEDEX_DB_PASSWORD"),
		os.Getenv("POKEDEX_DB_NAME"),
	)

	app.Run(":8989")

}
