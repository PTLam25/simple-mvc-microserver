package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNotFound(t *testing.T) {
	// 1) вызов функции для теста
	user, err := GetUser(0)

	// 2) валидация результата функции
	assert.Nil(t, user, "not expecting a user with id 0")
	assert.NotNilf(t, err, "expecting error when a user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "expecting 404 status code when a user with id 0")
	assert.EqualValues(t, "user 0 was not found", err.Message, "expecting 404 status code when a user with id 0")
	assert.EqualValues(t, "not_found", err.Code, "expecting 404 status code when a user with id 0")
}
