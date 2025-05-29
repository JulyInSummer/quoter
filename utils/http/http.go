package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// JSON writes a JSON http response with headers
// and status code 200
func JSON(w http.ResponseWriter, statusCode int, data any) {
	setHeaders(w)
	w.WriteHeader(statusCode)

	content := response{
		Code: statusCode,
		Data: data,
	}

	write(w, content)
}

// Handle checks whether the handler returned an error.
// If the handler returns an error, then status code 501 is written
func Handle(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		err := handler(w, r)
		if err != nil {
			HandleInternalError(w)
			return
		}
	}
}

func HandleValidationError(w http.ResponseWriter, errors []string) {
	setHeaders(w)
	w.WriteHeader(http.StatusUnprocessableEntity)

	content := map[string]any{
		"code":   http.StatusUnprocessableEntity,
		"data":   "Validation error",
		"errors": errors,
	}

	write(w, content)
}

// HandleInternalError writes status code 501
func HandleInternalError(w http.ResponseWriter) {
	setHeaders(w)
	w.WriteHeader(http.StatusInternalServerError)

	content := response{
		Code: http.StatusInternalServerError,
		Data: "Internal server error",
	}

	write(w, content)
}

// HandleBadRequest writes status code 400
func HandleBadRequest(w http.ResponseWriter, message string) {
	setHeaders(w)
	w.WriteHeader(http.StatusBadRequest)

	content := response{
		Code: http.StatusBadRequest,
		Data: message,
	}

	write(w, content)
}

// HandleNotFound writes status code 404
func HandleNotFound(w http.ResponseWriter, message string) {
	setHeaders(w)
	w.WriteHeader(http.StatusNotFound)

	content := response{
		Code: http.StatusNotFound,
		Data: message,
	}

	write(w, content)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func write(w http.ResponseWriter, data any) {
	setHeaders(w)
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to unmarshal content: %v, error: %v\n", data, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("failed to write response: %v, error: %v\n", bytes, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}
}
