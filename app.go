package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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
	// app.RegisterRoutes()
}

func (app *App) Run(addr string) {
	// run the app
	fmt.Println("Server listening on port ", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) RegisterRoutes() {}
