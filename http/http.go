package http

import (
	"encoding/json"
	"net/http"
)

// Error is a reusable, JSON serializable error message structure
type Error struct {
	Message string `json:"message"`
}

// ReplyJSON responds to an HTTP request with a JSON message
func ReplyJSON(writer http.ResponseWriter, response interface{}, code int) {
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(code)
}

// ReplyError responds to an HTTP request with an error message
func ReplyError(writer http.ResponseWriter, err error, code int) {
	ReplyJSON(writer, Error{Message: err.Error()}, code)
}

// ReplyInternalServerError responds to an HTTP request with an error message and 503 status
func ReplyInternalServerError(writer http.ResponseWriter, err error) {
	ReplyError(writer, err, http.StatusInternalServerError)
}
