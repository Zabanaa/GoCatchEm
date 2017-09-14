package models

type Pokemon struct {
	ID         int    `json:"id,omitempty"`
	Number     string `json:"number,omitempty"`
	Name       string `json:"name,omitempty"`
	JpName     string `json:"jp_name,omitempty"`
	Types      string `json:"types,omitempty"`
	Stats      Stats  `json:"stats,omitempty"`
	Bio        string `json:"bio,omitempty"`
	Generation int64  `json:"generation,omitempty"`
}

type Stats struct {
	HP      int `json:"hp,omitempty"`
	Attack  int `json:"attack,omitempty"`
	Defense int `json:"defense,omitempty"`
	Sp_atk  int `json:"sp_atk,omitempty"`
	Sp_def  int `json:"sp_def,omitempty"`
	Speed   int `json:"speed,omitempty"`
}
