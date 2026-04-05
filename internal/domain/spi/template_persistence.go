package spi

import "github.com/lian-corp/notification-microservice/internal/domain/model"

type TemplatePersistence interface {
	GetByEventType(eventType string) (model.NotificationTemplate, error)
}