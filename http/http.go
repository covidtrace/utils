package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAuthorization returns the `authType` authorization header, or an error
func GetAuthorization(r *http.Request, authType string) (string, error) {
	authorization := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(authorization) != 2 {
		return "", errors.New("missing authorization header")
	}

	if !strings.EqualFold(authorization[0], authType) {
		return "", fmt.Errorf("only %s authorization supported", authType)
	}

	return authorization[1], nil
}

// Error is a reusable, JSON serializable error message structure
type Error struct {
	Message string `json:"message"`
}

// ReplyJSON responds to an HTTP request with a JSON message
func ReplyJSON(writer http.ResponseWriter, response interface{}, code int) {
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

// ReplyError responds to an HTTP request with an error message
func ReplyError(writer http.ResponseWriter, err error, code int) {
	ReplyJSON(writer, Error{Message: err.Error()}, code)
}

// ReplyInternalServerError responds to an HTTP request with an error message and 503 status
func ReplyInternalServerError(writer http.ResponseWriter, err error) {
	ReplyError(writer, err, http.StatusInternalServerError)
}

// ReplyBadRequestError responds to an HTTP request with an error message and 400 status
func ReplyBadRequestError(writer http.ResponseWriter, err error) {
	ReplyError(writer, err, http.StatusBadRequest)
}
