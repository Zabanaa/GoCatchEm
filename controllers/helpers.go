package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Meta struct {
		Type       string `json:"type, omitempty"`
		StatusCode int64  `json:"status_code, omitempty"`
	}

	Body map[string]interface{} `json:"body, omitempty"`
}

func respondWithJson(w http.ResponseWriter, code int, payload Response) {

	response, _ := json.Marshal(payload)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}
