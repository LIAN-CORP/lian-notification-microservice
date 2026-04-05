package spi

import "github.com/lian-corp/notification-microservice/internal/domain/model"

type NotificationPersistence interface {
	Save(notification model.Notification) error
	Update(notification model.Notification) error
}