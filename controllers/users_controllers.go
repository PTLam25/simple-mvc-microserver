package controllers

import (
	"encoding/json"
	"github.com/PTLam25/microserver-course-1/services"
	"net/http"
	"strconv"
)

// Контроллеры - это входная точка для любого запроса в нашем приложения.
// Они отвечают за валидации входящих данных и передача их сервисам на обработку. А потом возвращают ответ клиенту.
// В контроллерах НЕ должна содеражаться бизнес логика. Бизнес логика содержатся в сервисах.

func GetUser(resp http.ResponseWriter, req *http.Request) {
	// 1) валидируем данные
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		// добавить 404 станус в заголовок ответа
		resp.WriteHeader(http.StatusBadRequest)
		// отправить текст ошибки
		resp.Write([]byte("user_id must be a number"))
	}

	// 2) отправляем данные на обработку сервису
	user, err := services.GetUser(userId)

	if err != nil {
		// добавить 404 станус в заголовок ответа
		resp.WriteHeader(http.StatusNotFound)
		// отправить текст ошибки
		resp.Write([]byte(err.Error()))
	}

	// 3) декодировать данные в JSON
	jsonValue, _ := json.Marshal(user)

	// 4) вернуть ответ клиенту
	resp.Write(jsonValue)
}
