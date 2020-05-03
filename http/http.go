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
func ReplyJSON(writer http.ResponseWriter, code int, response interface{}) {
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(code)
}

// ReplyError responds to an HTTP request with an error message
func ReplyError(writer http.ResponseWriter, code int, err error) {
	ReplyJSON(writer, code, Error{Message: err.Error()})
}
