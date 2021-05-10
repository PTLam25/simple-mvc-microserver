package domain

// DOA отвечает за связь модели С БД для CRUD.
import (
	"fmt"
	"github.com/PTLam25/microserver-course-1/utils"
	"log"
	"net/http"
)

type userDao struct {
}

var (
	// типо БД, в котором есть пользователи
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Lam", LastName: "Pham", Email: "test@gmail.com"},
	}
	UserDao userDao
)

func (ud *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We are accessing DB")
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
