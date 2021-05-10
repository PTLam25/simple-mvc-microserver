package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNotFound(t *testing.T) {
	// 1) вызов функции для теста
	user, err := GetUser(0)

	// 2) валидация результата функции
	if user != nil {
		t.Error("not expecting a user with id 0")
	}

	if err == nil {
		t.Error("expecting error when a user id is 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("expecting 404 status code when a user with id 0")
	}
}
