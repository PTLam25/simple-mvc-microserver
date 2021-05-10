package domain

import "errors"

// DOA отвечает за связь модели С БД для CRUD.

var (
	// типо БД, в котором есть пользователи
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Lam", LastName: "Pham", email: "test@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	//  получить данные с БД
	if user := users[userId]; user != nil {
		return user, nil

	}

	return nil, errors.New("NO SUCH USER")
}
