package http

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	rw := httptest.NewRecorder()

	content := "Success"

	JSON(rw, http.StatusCreated, content)

	assert.Equal(t, http.StatusCreated, rw.Result().StatusCode)
	assert.Equal(t, "application/json", rw.Header().Get("Content-Type"))

	var resp response
	err := json.Unmarshal(rw.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.Code)
	data, ok := resp.Data.(string)
	assert.True(t, ok)

	assert.Equal(t, "Success", data)
}

func TestHandleBadRequest(t *testing.T) {
	rw := httptest.NewRecorder()

	content := "Bad Request"

	HandleBadRequest(rw, content)

	assert.Equal(t, "application/json", rw.Header().Get("Content-Type"))

	var resp response
	err := json.Unmarshal(rw.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, content, resp.Data)
}

func TestHandleValidationError(t *testing.T) {
	rw := httptest.NewRecorder()

	content := []string{"ValidationError1", "ValidationError2"}

	HandleValidationError(rw, content)

	assert.Equal(t, "application/json", rw.Header().Get("Content-Type"))

	var resp validationResponse
	err := json.Unmarshal(rw.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
	assert.Equal(t, "Validation error", resp.Data)

	assert.EqualValues(t, content, resp.Errors)
}

func TestHandleInternalServerError(t *testing.T) {
	rw := httptest.NewRecorder()

	HandleInternalError(rw)

	assert.Equal(t, "application/json", rw.Header().Get("Content-Type"))

	var resp response
	err := json.Unmarshal(rw.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Equal(t, "Internal Server Error", resp.Data)
}

func TestHandleNotFound(t *testing.T) {
	rw := httptest.NewRecorder()

	content := "Requested Resource Not Found"

	HandleNotFound(rw, content)

	assert.Equal(t, "application/json", rw.Header().Get("Content-Type"))

	var resp response
	err := json.Unmarshal(rw.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Equal(t, content, resp.Data)
}
