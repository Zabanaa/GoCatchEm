package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"pokemon_api/controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(dbuser, dbpassword, dbname string) {

	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", dbuser, dbpassword, dbname)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.RegisterRoutes()
}

func (app *App) Run(addr string) {
	// run the app
	fmt.Println("Server listening on port ", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) RegisterRoutes() {
	app.Router.HandleFunc("/pokemons", app.getAllPokemons).Methods("GET")
	app.Router.HandleFunc("/pokemons", app.createPokemon).Methods("POST")
	app.Router.HandleFunc("/pokemons/{name}", app.getPokemon).Methods("GET")
	app.Router.HandleFunc("/pokemons/{name}", app.deletePokemon).Methods("DELETE")
}

// Controllers

func (app *App) getAllPokemons(w http.ResponseWriter, r *http.Request) {
	controllers.GetAllPokemons(app.DB, w, r)
}

func (app *App) getPokemon(w http.ResponseWriter, r *http.Request) {
	controllers.GetPokemon(app.DB, w, r)
}

func (app *App) deletePokemon(w http.ResponseWriter, r *http.Request) {
	controllers.DeletePokemon(app.DB, w, r)
}

func (app *App) createPokemon(w http.ResponseWriter, r *http.Request) {
	controllers.CreatePokemon(app.DB, w, r)
}
