package spi

import "github.com/lian-corp/notification-microservice/internal/domain/model"

type NotificationEventPersistence interface {
	Save(event model.NotificationEvent) error
	MarkAsProcessed(eventID string) error
}