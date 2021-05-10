package domain

// модель пользователя в нашем приложения

type User struct {
	Id        uint64
	FirstName string
	LastName  string
	email     string
}
