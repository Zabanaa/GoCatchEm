package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS pokemons (
	id SERIAL PRIMARY KEY,
	number CHAR(10) NOT NULL,
	name CHAR(255) NOT NULL,
	jp_name CHAR(255) NOT NULL,
	types CHAR(255) NOT NULL,
	hp INT NOT NULL,
	attack INT NOT NULL,
	defense INT NOT NULL,
	sp_atk INT NOT NULL,
	sp_def INT NOT NULL,
	speed INT NOT NULL,
	bio TEXT NOT NULL,
	generation INT,
	CONSTRAINT unique_name UNIQUE (name)
);
`

var app App

func TestMain(m *testing.M) {

	app = App{}

	app.Initialize(
		os.Getenv("TEST_POKEDEX_DB_USER"),
		os.Getenv("TEST_POKEDEX_DB_PASSWORD"),
		os.Getenv("TEST_POKEDEX_DB_NAME"),
	)

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/pokemons", nil)
	response := executeRequest(req)

	var jsonResponse map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &jsonResponse)

	pokemonCount := jsonResponse["body"].(map[string]interface{})["count"].(float64)

	if pokemonCount != 0 {
		t.Errorf("Expected data to be null, got %s", jsonResponse)
	}

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestNonExistentPokemon(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/pokemons/pikachu", nil)
	response := executeRequest(req)

	var jsonResponse map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &jsonResponse)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	errMessage := jsonResponse["body"].(map[string]interface{})["message"]
	expectedErrMessage := "This pokemon does not exist"

	if errMessage != expectedErrMessage {
		t.Errorf("Expected %s got %s", expectedErrMessage, errMessage)
	}
}

func ensureTableExists() {

	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM pokemons;")
	app.DB.Exec("ALTER SEQUENCE pokemons_id_seq RESTART WITH 1;")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code: %d. Got %d", expected, actual)
	}
}
