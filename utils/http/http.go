package http

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON writes a JSON http response with headers
// and status code 200
func JSON(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusOK)

	content := map[string]any{
		"code":    200,
		"message": data,
	}

	bytes, err := json.Marshal(content)
	if err != nil {
		log.Printf("failed to unmarshal content: %v, error: %v\n", content, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("failed to write response: %v, error: %v\n", bytes, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

}

// Handle checks whether the handler returned an error.
// If the handler returns an error, then status code 501 is written
func Handle(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := handler(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
	}
}

func HandleValidationError(w http.ResponseWriter, errors []string) {
	w.WriteHeader(http.StatusUnprocessableEntity)

	content := map[string]any{
		"code":    http.StatusUnprocessableEntity,
		"message": "Validation error",
		"errors":  errors,
	}

	bytes, err := json.Marshal(content)
	if err != nil {
		log.Printf("failed to unmarshal content: %v, error: %v\n", content, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("failed to write response: %v, error: %v\n", bytes, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
}

// HandleInternalError writes status code 501
func HandleInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error"))
	return
}

// HandleBadRequest writes status code 400
func HandleBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request"))
	return
}
