package services

import "github.com/PTLam25/microserver-course-1/domain"

// Сервисы содержат в себя бизнес логику для обработки данных от контроллеров. Сервисы могут общаться с другими сервисами (вызывать у себя в коде), а также к DAO для связи с БД, для получения данных.

func GetUser(userId int64) (*domain.User, error) {
	// тут может быть вызовы к другим сервисам
	return domain.GetUser(userId)
}
