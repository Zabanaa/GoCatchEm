package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"pokemon_api/app/controllers"

	"github.com/gorilla/handlers"
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

	if err := app.DB.Ping(); err != nil {
		panic(err)
	}

	app.Router = mux.NewRouter()
	app.RegisterRoutes()
}

func (app *App) Run(addr string) {
	// run the app

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Server listening on port ", addr)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(headersOk, originsOk, methodsOk)(app.Router)))
}

func (app *App) RegisterRoutes() {
	app.Router.HandleFunc("/pokemons", app.getAllPokemons).Methods("GET")
	app.Router.HandleFunc("/pokemons/type/{type}", app.getPokemonsByType).Methods("GET")
	app.Router.HandleFunc("/pokemons/generation/{generation}", app.getPokemonsByGen).Methods("GET")
	app.Router.HandleFunc("/pokemons", app.createPokemon).Methods("POST")
	app.Router.HandleFunc("/pokemons/{name}", app.getPokemon).Methods("GET")
	app.Router.HandleFunc("/pokemons/{name}", app.updatePokemon).Methods("PUT")
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

func (app *App) updatePokemon(w http.ResponseWriter, r *http.Request) {
	controllers.UpdatePokemon(app.DB, w, r)
}

func (app *App) getPokemonsByType(w http.ResponseWriter, r *http.Request) {
	controllers.GetPokemonsByType(app.DB, w, r)
}

func (app *App) getPokemonsByGen(w http.ResponseWriter, r *http.Request) {
	controllers.GetPokemonsByGen(app.DB, w, r)
}
