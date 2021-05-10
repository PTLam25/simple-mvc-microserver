package controllers

// Контроллеры - это входная точка для любого запроса в нашем приложения.
// Они отвечают за валидации входящих данных и передача их сервисам на обработку. А потом возвращают ответ клиенту.
// В контроллерах НЕ должна содеражаться бизнес логика. Бизнес логика содержатся в сервисах.
import (
	"encoding/json"
	"fmt"
	"github.com/PTLam25/microserver-course-1/services"
	"github.com/PTLam25/microserver-course-1/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	// 1) валидируем данные
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "bad_request",
		}
		jsonErrorValue, _ := json.Marshal(apiErr)
		// добавить 404 станус в заголовок ответа
		resp.WriteHeader(apiErr.StatusCode)
		// отправить текст ошибки
		resp.Write(jsonErrorValue)
		return
	}

	// 2) отправляем данные на обработку сервису
	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		jsonErrorValue, _ := json.Marshal(apiErr)
		// добавить 404 станус в заголовок ответа
		resp.WriteHeader(apiErr.StatusCode)
		// отправить текст ошибки
		resp.Write(jsonErrorValue)
		return
	}

	// 3) декодировать данные в JSON
	jsonValue, _ := json.Marshal(user)

	// 4) вернуть ответ клиенту
	resp.Write(jsonValue)
}
