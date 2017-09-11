package models

import (
	"database/sql"
	"errors"
)

type Pokemon struct {
	ID         int    `json:"id:omitempty"`
	Number     string `json:"number:omitempty"`
	Name       string `json:"name:omitempty"`
	JpName     string `json:"jp_name:omitempty"`
	Types      string `json:"types:omitempty"`
	Stats      Stats  `json:"stats:omitempty"`
	Bio        string `json:"bio:omitempty"`
	Generation int    `json:"generation:omitempty"`
}

type Stats struct {
	HP      int `json:"hp:omitempty"`
	Attack  int `json:"attack:omitempty"`
	Defense int `json:"defense:omitempty"`
	Sp_atk  int `json:"sp_atk:omitempty"`
	Sp_def  int `json:"sp_def:omitempty"`
	Speed   int `json:"speed:omitempty"`
}

func GetPokemons(db *sql.DB) ([]Pokemon, error) {

	query := "SELECT * FROM pokemons"

	// Perform the select query
	rows, err := db.Query(query)

	// If it fails return nil and and the error
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pokemons []Pokemon

	for rows.Next() {
		var stats Stats
		var pokemon Pokemon

		err := rows.Scan(
			&pokemon.ID, &pokemon.Number, &pokemon.Name, &pokemon.JpName,
			&pokemon.Types, &stats.HP, &stats.Attack, &stats.Defense,
			&stats.Sp_atk, &stats.Sp_def, &stats.Speed, &pokemon.Bio,
			&pokemon.Generation)

		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

// func createPokemon(db *sql.DB) (*Pokemon, error) {
// 	return errors.New("Not Implemented")
// }

func (pokemon *Pokemon) getInfo(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (pokemon *Pokemon) updateInfo(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (pokemon *Pokemon) deleteInfo(db *sql.DB) error {
	return errors.New("Not Implemented")
}
