package models

type Pokemon struct {
	ID         int    `json:"id, omitempty"`
	Number     string `json:"number"`
	Name       string `json:"name"`
	JpName     string `json:"jp_name"`
	Types      string `json:"types"`
	Stats      Stats  `json:"stats"`
	Bio        string `json:"bio"`
	Generation int64  `json:"generation"`
}

type Stats struct {
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	Sp_atk  int `json:"sp_atk"`
	Sp_def  int `json:"sp_def"`
	Speed   int `json:"speed"`
}
