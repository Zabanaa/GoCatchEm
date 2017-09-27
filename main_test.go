package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const pikachu = `{
	"number":  "#025",
	"name": "pikachu",
	"jp_name": "pikachu",
	"types": "electric",
	"stats": {
		"hp": 23,
		"attack": 24,
		"defense": 93,
		"sp_atk": 23,
		"sp_def": 88,
		"speed": 12
	},
	"bio": "Pika pika",
	"generation": 1
}`

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

func TestCreatePokemon(t *testing.T) {
	clearTable()
	code, pikachu := createPikachu()

	checkResponseCode(t, http.StatusCreated, code)

	message := pikachu["body"].(map[string]interface{})["message"]
	expectedMessage := "Pokemon Created"

	if message != expectedMessage {
		t.Errorf("Expected message: %s. Got: '%s' instead", expectedMessage, message)
	}
}

func TestDuplicatePokemon(t *testing.T) {
	clearTable()
	createPikachu()
	code, pikachu := createPikachu()

	checkResponseCode(t, http.StatusConflict, code)

	errMessage := pikachu["body"].(map[string]interface{})["message"].(string)
	expectedMessage := "A pokemon with this name already exists"

	if !strings.Contains(errMessage, expectedMessage) {
		t.Errorf("Expected message: %s. Got: '%s' instead", expectedMessage, errMessage)
	}
}

func TestMissingFields(t *testing.T) {

	clearTable()
	payload := []byte(`{ "number": "#025", "name": "raichu",
	"types": "electric", "stats": { "hp": 23,
	"defense": 93, "sp_atk": 23, "sp_def": 88, "speed": 12 },
	"bio": "Pika pika", "generation": 1 }`)

	code, raichu := createPokemon(payload)
	checkResponseCode(t, http.StatusUnprocessableEntity, code)

	errMessage := raichu["body"].(map[string]interface{})["message"].(string)
	expectedErrMessage := "Missing fields"

	if expectedErrMessage != errMessage {
		t.Errorf("Expected message: %s. Got: '%s' instead", expectedErrMessage, errMessage)
	}
}

func TestUpdatePokemon(t *testing.T) {

	clearTable()
	createPikachu()

	payload := []byte(`{
		"number":  "#025",
		"name": "pikachuuuu",
		"jp_name": "pikachu",
		"types": "electric",
		"stats": {
			"hp": 23,
			"attack": 24,
			"defense": 93,
			"sp_atk": 23,
			"sp_def": 88,
			"speed": 12
		},
		"bio": "Pika pika",
		"generation": 1
	}`)

	code, updatedPikachu := updatePokemon("pikachu", payload)

	checkResponseCode(t, http.StatusOK, code)

	updatedName := updatedPikachu["body"].(map[string]interface{})["pokemon"]

	updatedName = updatedName.(map[string]interface{})["name"].(string)
	expectedName := "pikachuuuu"

	if updatedName != expectedName {
		t.Errorf("Expected name: %s. Got: '%s' instead", updatedName, expectedName)
	}
}

func TestUpdateNonExistentPokemon(t *testing.T) {

	clearTable()
	createPikachu()

	payload := []byte(`{ "name": "pikachuuuu" }`)

	code, updatedPikachu := updatePokemon("karimbenzema", payload)

	checkResponseCode(t, http.StatusNotFound, code)

	expectedErrMessage := "This Pokemon does not exist"
	errMessage := updatedPikachu["body"].(map[string]interface{})["message"].(string)

	if errMessage != expectedErrMessage {
		t.Errorf("Expected name: %s. Got: '%s' instead", expectedErrMessage, errMessage)
	}
}

func TestDeletePokemon(t *testing.T) {

	clearTable()
	createPikachu()

	code, _ := deletePikachu()

	checkResponseCode(t, http.StatusNoContent, code)
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

func updatePokemon(name string, payload []byte) (int, map[string]interface{}) {

	req, _ := http.NewRequest("PUT", "/pokemons/"+name, bytes.NewBuffer(payload))
	response := executeRequest(req)

	var jsonResponse map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	return response.Code, jsonResponse

}

func createPikachu() (int, map[string]interface{}) {
	req, _ := http.NewRequest("POST", "/pokemons", bytes.NewBuffer([]byte(pikachu)))
	response := executeRequest(req)

	var jsonResponse map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	return response.Code, jsonResponse
}

func createPokemon(payload []byte) (int, map[string]interface{}) {
	req, _ := http.NewRequest("POST", "/pokemons", bytes.NewBuffer([]byte(payload)))
	response := executeRequest(req)

	var jsonResponse map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	return response.Code, jsonResponse
}

func deletePikachu() (int, map[string]interface{}) {

	req, _ := http.NewRequest("DELETE", "/pokemons/pikachu", nil)
	response := executeRequest(req)

	var jsonResponse map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	return response.Code, jsonResponse
}
