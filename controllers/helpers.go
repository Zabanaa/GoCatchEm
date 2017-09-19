package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

type Response struct {
	Meta Meta                   `json:"meta,omitempty"`
	Body map[string]interface{} `json:"body, omitempty"`
}

type Meta struct {
	Type       string `json:"type, omitempty"`
	StatusCode int64  `json:"status_code, omitempty"`
}

func (response *Response) StatusOK(w http.ResponseWriter, body map[string]interface{}) {

	response.Meta.Type = "success"
	response.Meta.StatusCode = http.StatusOK
	response.Body = body
	respondWithJson(w, response.Meta.StatusCode, response)
}

func (response *Response) MissingFields(w http.ResponseWriter, missingField string) {

	response.Meta.Type = "error"
	response.Meta.StatusCode = http.StatusUnprocessableEntity
	response.Body = map[string]interface{}{
		"message":        "Missing fields",
		"missing_fields": missingField,
	}
	respondWithJson(w, response.Meta.StatusCode, response)
}

// This Helper should really accept a field name
// To handle the case where I add another unique constraint to the pokemon table
// but it will do for now
func (response *Response) Conflict(w http.ResponseWriter) {

	response.Meta.Type = "error"
	response.Meta.StatusCode = http.StatusConflict
	response.Body = map[string]interface{}{
		"message": "A pokemon with this name already exists",
	}
	respondWithJson(w, response.Meta.StatusCode, response)
}

func (response *Response) BadRequest(w http.ResponseWriter) {

	response.Meta.Type = "error"
	response.Meta.StatusCode = http.StatusBadRequest
	response.Body = map[string]interface{}{
		"message": "Couldn't process the request. Make sure it's properly formatted and that the fields are of the correct types. For more information on types, please refer to the documentation.",
	}
	respondWithJson(w, response.Meta.StatusCode, response)
}

func (response *Response) NotFound(w http.ResponseWriter) {

	response.Meta.Type = "error"
	response.Meta.StatusCode = http.StatusNotFound
	response.Body = map[string]interface{}{
		"message": "Pokemon not found.",
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

func extractMissingField(errorMessage string) string {
	pattern := regexp.MustCompile(`\"(\w+)\"`)
	missingField := strings.Replace(pattern.FindString(errorMessage), "\"", "", 2)
	return missingField
}
