package app

// отвечает за конфигурацию и запуск севера

import (
	"github.com/PTLam25/microserver-course-1/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
