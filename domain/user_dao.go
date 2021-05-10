package domain

import (
	"fmt"
	"github.com/PTLam25/microserver-course-1/utils"
	"net/http"
)

// DOA отвечает за связь модели С БД для CRUD.

var (
	// типо БД, в котором есть пользователи
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Lam", LastName: "Pham", Email: "test@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	//  получить данные с БД
	if user := users[userId]; user != nil {
		return user, nil

	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
