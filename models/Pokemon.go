package models

type Pokemon struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Types      string `json:"types"`
	Stats      Stats  `json:"stats"`
	Bio        string `json:"bio"`
	Generation int    `json:"generation"`
}

type Stats struct {
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	Sp_atk  int `json:"sp_atk"`
	Sp_def  int `json:"sp_def"`
	Speed   int `json:"speed"`
}

func getProducts(db *sql.DB) ([]Pokemon, error) {

}

func createPokemon(db *sql.DB) (*Pokemon, error) {

}

func (pokemon *Pokemon) getInfo(db *sql.DB) error {

}

func (pokemon *Pokemon) updateInfo(db *sql.DB) error {}
func (pokemon *Pokemon) deleteInfo(db *sql.DB) error {}
