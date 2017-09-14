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

func (response *Response) StatusOK(w http.ResponseWriter, body map[string]interface{}) {

	response.Meta.Type = "success"
	response.Meta.StatusCode = http.StatusOK
	response.Body = body
	respondWithJson(w, response.Meta.StatusCode, response)
}

func (response *Response) BadRequest(w http.ResponseWriter, message string) {

	response.Meta.Type = "error"
	response.Meta.StatusCode = http.StatusBadRequest
	response.Body = map[string]interface{}{
		"message": message,
	}
	respondWithJson(w, response.Meta.StatusCode, response)
}

func (response *Response) NotFound(w http.ResponseWriter, message string) {

	response.Meta.Type = "error"
	response.Meta.StatusCode = http.StatusNotFound
	response.Body = map[string]interface{}{
		"message": message,
	}
	respondWithJson(w, response.Meta.StatusCode, response)
}

func (response *Response) ServerError(w http.ResponseWriter, message string) {

	response.Meta.Type = "error"
	response.Meta.StatusCode = http.StatusInternalServerError
	response.Body = map[string]interface{}{
		"message": message,
	}
	respondWithJson(w, response.Meta.StatusCode, response)
}

func (response *Response) Created(w http.ResponseWriter, message string) {

	response.Meta.Type = "success"
	response.Meta.StatusCode = http.StatusCreated
	response.Body = map[string]interface{}{
		"message": message,
	}
	respondWithJson(w, response.Meta.StatusCode, response)
}

func (response *Response) Deleted(w http.ResponseWriter) {

	response.Meta.Type = "success"
	response.Meta.StatusCode = http.StatusNoContent
	respondWithJson(w, response.Meta.StatusCode, response)
}

func respondWithJson(w http.ResponseWriter, code int64, payload *Response) {

	response, _ := json.Marshal(payload)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(int(code)) // We convert code to an int because w.WriteHeader does not accept Int64 type values
	w.Write(response)
}
